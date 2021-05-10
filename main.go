package main

import (
	"github.com/thecodingmachine/gotenberg-go-client/v7"
	"github.com/yigitsadic/hsedocumentgenerator/internal/compressor"
	"github.com/yigitsadic/hsedocumentgenerator/internal/handlers"
	"github.com/yigitsadic/hsedocumentgenerator/internal/models"
	"github.com/yigitsadic/hsedocumentgenerator/internal/pdf_generator"
	"os"
)

type mockGoogleClient struct {
}

func (m mockGoogleClient) ReadFromSheets() ([]models.Record, error) {
	return []models.Record{
		{
			FirstName:       "Yiğit",
			LastName:        "Sadıç",
			JobName:         "Forklift Şoförü",
			CompanyName:     "DAL Heavy",
			EducationDate:   "21.05.2021",
			EducationName:   "Güvenli Sürüş",
			EducationHours:  "22 saatlik",
			UniqueReference: "abE1Ec1-A",
		},
	}, nil
}

func main() {
	gotenbergUrl := os.Getenv("GOTENBERG_URL")
	if gotenbergUrl == "" {
		gotenbergUrl = "http://localhost:3000"
	}

	assetStore := pdf_generator.NewStore()

	googleClient := mockGoogleClient{}
	pdfGen := &pdf_generator.PDFGenerator{
		Store:           assetStore,
		GotenbergClient: gotenberg.Client{Hostname: gotenbergUrl},
	}
	z := compressor.Zipper{}

	h := handlers.NewHandler(os.Stdin, os.Stdout, googleClient, pdfGen, z)

	h.Do()
}
