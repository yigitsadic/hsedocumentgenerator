package handlers

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"github.com/yigitsadic/hsedocumentgenerator/internal/compressor"
	"github.com/yigitsadic/hsedocumentgenerator/internal/models"
	"github.com/yigitsadic/hsedocumentgenerator/internal/pdf_generator"
	"github.com/yigitsadic/hsedocumentgenerator/internal/sheet_reader"
	"io"
	"log"
	"strings"
)

const (
	welcomeText                        = "🚀\tGoogle Sheets üzerinden okuma başlatıldı.\n"
	recordReadText                     = "📗\tGoogle Sheets üzerinden %d kayıt okundu.\n"
	cannotReadFromGoogleText           = "😥\tGoogle Sheets üzerinden kayıtlar okunamadı.\n"
	outputZIPText                      = "🤔\tOluşturulan PDFleri içeren ZIP dosyası nereye depolanacak?:\t"
	pdfGenerationStartedText           = "⏳\tPDF belge üretme işlemi başlandı...\n"
	pdfGeneratedText                   = "👍\t[%s.pdf] %s %s için PDF belgesi üretildi.\n"
	zipFileCreatedText                 = "✅\tPDF belgeleri %q olarak sıkıştırıldı ve okunan kayıtlar Google Sheets içine eklendi.\n"
	processSucceededText               = "💫\tİşlem tamamlandı. İyi günler!\n"
	noRecordFoundText                  = "\U0001F97A\tGoogle Sheets üzerinde kayıt bulunamadı. Yapacak bir şey yok.\n"
	errorOccurredDuringPDFCreationText = "😥\t[%s.pdf] %s %s için beklenmedik bir hata oluştu.\n"
	noFileToCompressText               = "🙈\tSıkıştırılacak PDF bulunamadı.\n"
)

type Handler struct {
	Output      io.Writer
	Reader      *bufio.Reader
	Client      sheet_reader.SheetClient
	ReadRecords []models.Record

	PDFGenerator pdf_generator.PDFGenerate

	Files []models.ReadFile

	ZipWriter     compressor.ZipWriter
	ZipOutputPath string
}

func NewHandler(input io.Reader,
	output io.Writer,
	client sheet_reader.SheetClient,
	pdfGen pdf_generator.PDFGenerate,
	zipper compressor.ZipWriter,
) *Handler {
	return &Handler{
		Output:       output,
		Reader:       bufio.NewReader(input),
		Client:       client,
		ReadRecords:  []models.Record{},
		Files:        []models.ReadFile{},
		PDFGenerator: pdfGen,
		ZipWriter:    zipper,
	}
}

// Writes to output source
func (h *Handler) Write(text string, inputs ...interface{}) {
	if len(inputs) == 0 {
		fmt.Fprint(h.Output, text)
	} else {
		fmt.Fprintf(h.Output, text, inputs...)
	}
}

// Reads from Google Sheet client and prints output
func (h *Handler) ReadFromSheets() error {
	result, err := h.Client.ReadFromSheets()
	if err != nil {
		fmt.Fprint(h.Output, cannotReadFromGoogleText)

		return err
	}

	h.ReadRecords = result

	fmt.Fprintf(h.Output, fmt.Sprintf(recordReadText, len(h.ReadRecords)))
	return nil
}

// Reads output path and stores it.
func (h *Handler) StoreOutputPath() {
	fmt.Fprint(h.Output, outputZIPText)

	text, _ := h.Reader.ReadString('\n')

	h.ZipOutputPath = strings.TrimSpace(text)
}

// Generates PDF and stores it.
func (h *Handler) GeneratePDF(r models.Record) error {
	req, err := h.PDFGenerator.BuildRequest(r)
	if err != nil {
		return err
	}

	result, err := h.PDFGenerator.Build(req)
	if err != nil {
		return err
	}

	h.Files = append(h.Files, models.ReadFile{FileName: fmt.Sprintf("%s.pdf", r.UniqueReference), Content: result})

	h.Write(pdfGeneratedText, r.UniqueReference, r.FirstName, r.LastName)
	return nil
}

// Compresses PDF files and writes them as zip file.
func (h *Handler) WriteFilesToZip() error {
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	for _, file := range h.Files {
		f, err := w.Create(file.FileName)
		if err != nil {
			log.Println(err)
		}
		_, err = f.Write(file.Content)
		if err != nil {
			log.Println(err)
		}
	}

	err := w.Close()
	if err != nil {
		return err
	}

	err = h.ZipWriter.WriteAsZip(buf.Bytes())
	if err != nil {
		return err
	}

	h.Write(zipFileCreatedText, h.ZipOutputPath)
	return nil
}

// Handles all work flow.
func (h *Handler) Do() {
	h.Write(welcomeText)
	h.ReadFromSheets()

	if len(h.ReadRecords) == 0 {
		h.Write(noRecordFoundText)
		return
	}

	h.StoreOutputPath()

	h.Write(pdfGenerationStartedText)

	for _, record := range h.ReadRecords {
		err := h.GeneratePDF(record)
		if err != nil {
			h.Write(errorOccurredDuringPDFCreationText, record.UniqueReference, record.FirstName, record.LastName)
		}
	}

	if len(h.Files) == 0 {
		h.Write(noFileToCompressText)
		return
	} else {
		h.WriteFilesToZip()
	}

	h.Write(processSucceededText)
}
