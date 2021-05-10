package handlers

import (
	"bufio"
	"fmt"
	"github.com/yigitsadic/hsedocumentgenerator/internal/models"
	"github.com/yigitsadic/hsedocumentgenerator/internal/sheet_reader"
	"io"
)

const (
	welcomeText              = "🚀 Google Sheets üzerinden okuma başlatıldı.\n"
	recordReadText           = "📗 Google Sheets üzerinden %d kayıt okundu.\n"
	cannotReadFromGoogleText = "😥 Google Sheets üzerinden kayıtlar okunamadı.\n"
	outputZIPText            = "🤔 Oluşturulan PDFleri içeren ZIP dosyası nereye depolanacak?:"
	pdfGenerationStartedText = "⏳ PDF belge üretme işlemi başlandı...\n"
	pdfGeneratedText         = "👍 [%d/%d] PDF belgesi üretildi.\n"
	zipFileCreatedText       = "✅ PDF belgeleri %q olarak sıkıştırıldı ve Google Sheets içine eklendi.\n"
	processSucceededText     = "💫 İşlem tamamlandı. İyi günler!\n"
)

type Handler struct {
	Output      io.Writer
	Reader      *bufio.Reader
	Client      sheet_reader.SheetClient
	ReadRecords []models.Record

	ZipOutputPath string
}

func NewHandler(input io.Reader, output io.Writer, client sheet_reader.SheetClient) *Handler {
	return &Handler{
		Output:      output,
		Reader:      bufio.NewReader(input),
		Client:      client,
		ReadRecords: []models.Record{},
	}
}

// Greets user.
func (h *Handler) PrintHelloText() {
	fmt.Fprint(h.Output, welcomeText)
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

	h.ZipOutputPath = text
}

// Prints PDF generation process started text.
func (h *Handler) PrintPDFGenerationStarted() {
	fmt.Fprint(h.Output, pdfGenerationStartedText)
}
