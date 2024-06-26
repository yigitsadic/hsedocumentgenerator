package main

import (
	"github.com/rakyll/statik/fs"
	"github.com/thecodingmachine/gotenberg-go-client/v7"
	"github.com/yigitsadic/hsedocumentgenerator/internal/compressor"
	"github.com/yigitsadic/hsedocumentgenerator/internal/handlers"
	"github.com/yigitsadic/hsedocumentgenerator/internal/pdf_generator"
	"github.com/yigitsadic/hsedocumentgenerator/internal/sheet_handler"
	"log"
	"os"

	_ "github.com/yigitsadic/hsedocumentgenerator/statik"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	gotenbergUrl := os.Getenv("GOTENBERG_URL")
	if gotenbergUrl == "" {
		gotenbergUrl = "http://localhost:3000"
	}

	assetStore := pdf_generator.NewStore(statikFS)

	googleClient := sheet_handler.SheetHandler{}

	pdfGen := &pdf_generator.PDFGenerator{
		Store:           assetStore,
		GotenbergClient: gotenberg.Client{Hostname: gotenbergUrl},
	}

	if err = pdfGen.Ping(); err != nil {
		log.Fatal(err)
	}

	z := compressor.Zipper{}

	programHandler := handlers.NewHandler(os.Stdin, os.Stdout, googleClient, pdfGen, z)
	programHandler.Do()
}
