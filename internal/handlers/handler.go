package handlers

import (
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
	Client      sheet_reader.SheetClient
	ReadRecords []models.Record

	ZipOutputPath string
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
