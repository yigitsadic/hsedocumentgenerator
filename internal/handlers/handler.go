package handlers

import (
	"fmt"
	"github.com/yigitsadic/hsedocumentgenerator/internal/sheet_reader"
	"io"
)

const (
	welcomeText              = "🚀 Google Sheets üzerinden okuma başlatıldı."
	recordReadText           = "📗 Google Sheets üzerinden %d kayıt okundu."
	outputZIPText            = "🤔 Oluşturulan PDFleri içeren ZIP dosyası nereye depolanacak?:"
	pdfGenerationStartedText = "⏳ PDF belge üretme işlemi başlandı..."
	pdfGeneratedText         = "👍 [%d/%d] PDF belgesi üretildi."
	zipFileCreatedText       = "✅ PDF belgeleri %q olarak sıkıştırıldı ve Google Sheets içine eklendi."
	processSucceededText     = "💫 İşlem tamamlandı. İyi günler!"
)

type Handler struct {
	Output io.Writer
	Client *sheet_reader.SheetClient
}

// Greets user.
func (h *Handler) PrintHelloText() {
	fmt.Fprintln(h.Output, welcomeText)
}
