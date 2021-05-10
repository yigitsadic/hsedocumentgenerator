package main

import (
	"github.com/thecodingmachine/gotenberg-go-client/v7"
	"github.com/yigitsadic/hsedocumentgenerator/internal/compressor"
	"github.com/yigitsadic/hsedocumentgenerator/internal/handlers"
	"github.com/yigitsadic/hsedocumentgenerator/internal/pdf_generator"
	"github.com/yigitsadic/hsedocumentgenerator/internal/sheet_handler"
	"os"
)

func main() {
	gotenbergUrl := os.Getenv("GOTENBERG_URL")
	if gotenbergUrl == "" {
		gotenbergUrl = "http://localhost:3000"
	}

	assetStore := pdf_generator.NewStore()

	googleClient := sheet_handler.SheetHandler{}

	pdfGen := &pdf_generator.PDFGenerator{
		Store:           assetStore,
		GotenbergClient: gotenberg.Client{Hostname: gotenbergUrl},
	}
	z := compressor.Zipper{}

	programHandler := handlers.NewHandler(os.Stdin, os.Stdout, googleClient, pdfGen, z)
	programHandler.Do()
}
