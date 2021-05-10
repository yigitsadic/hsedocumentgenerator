package pdf_generator

import (
	"html/template"
	"io"
	"os"
)

type AssetStore struct {
	StateLogo        []byte
	Background       []byte
	HseLogo          []byte
	CompanySignature []byte
	Styles           []byte
	Template         *template.Template
}

// Reads all required files and stores them as bytes.
func NewStore() *AssetStore {
	s := AssetStore{}

	// Open and read state logo
	f1, err := os.Open("./assets/1920px-Emblem_of_Uzbekistan.svg.png")
	if err != nil {
		panic("Cannot open 1920px-Emblem_of_Uzbekistan.svg.png")
	}

	stateLogo, err := io.ReadAll(f1)
	if err != nil {
		panic("Cannot read 1920px-Emblem_of_Uzbekistan.svg.png")
	}
	f1.Close()

	s.StateLogo = stateLogo

	// Open and read background image
	f2, err := os.Open("./assets/bg.jpg")
	if err != nil {
		panic("Cannot open bg.jpg")
	}

	backgroundImage, err := io.ReadAll(f2)
	if err != nil {
		panic("Cannot read bg.jpg")
	}
	f2.Close()

	s.Background = backgroundImage

	// Open and read HSE Logo
	f3, err := os.Open("./assets/hse_logo.png")
	if err != nil {
		panic("Cannot open hse_logo.png")
	}

	hseLogo, err := io.ReadAll(f3)
	if err != nil {
		panic("Cannot read hse_logo.png")
	}
	f3.Close()

	s.HseLogo = hseLogo

	// Open and read company signature image
	f4, err := os.Open("./assets/kase.png")
	if err != nil {
		panic("Cannot open kase.png")
	}

	companySignature, err := io.ReadAll(f4)
	if err != nil {
		panic("Cannot read kase.png")
	}
	f4.Close()

	s.CompanySignature = companySignature

	// Open and read CSS file
	f5, err := os.Open("./assets/styles.css")
	if err != nil {
		panic("Cannot open styles.css")
	}

	styles, err := io.ReadAll(f5)
	if err != nil {
		panic("Cannot read styles.css")
	}
	f5.Close()

	s.Styles = styles

	// Open and read HTML -> PDF template file
	t, err := template.ParseFiles("./assets/certificate_template.html")
	if err != nil {
		panic("Cannot open certificate_template.html")
	}

	s.Template = t

	return &s
}
