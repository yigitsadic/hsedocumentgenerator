package models

import (
	"errors"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/skip2/go-qrcode"
	"log"
)

const (
	charset = "123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-"
	length  = 17

	baseURL = "https://hsegroup.uz/kurumsal/certificate_verification?qr_code="
)

var (
	UniqueReferenceMustPresentErr = errors.New("unique reference must present")
)

type Record struct {
	FirstName       string
	LastName        string
	JobName         string
	CompanyName     string
	EducationDate   string
	EducationName   string
	EducationHours  string
	UniqueReference string
}

// Assigns unique nano id reference code to object.
func (r *Record) AssignUniqueReference() {
	code, err := gonanoid.Generate(charset, length)
	if err != nil {
		log.Fatalln("Error occurred while assigning unique reference")

		return
	}

	r.UniqueReference = code
}

// Creates QR code for record and returns it as bytes.
func (r Record) GenerateQRCode() ([]byte, error) {
	if r.UniqueReference == "" {
		return nil, UniqueReferenceMustPresentErr
	}

	return qrcode.Encode(baseURL+r.UniqueReference, qrcode.Medium, 100)
}
