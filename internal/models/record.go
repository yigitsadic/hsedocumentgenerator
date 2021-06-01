package models

import (
	"errors"
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/skip2/go-qrcode"
	"time"
)

const (
	charset = "123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-"
	length  = 17

	trBaseURL = "https://hsegroup.uz/kurumsal/sertifika-dogrulama/?qr_code=%s"
	enBaseURL = "https://hsegroup.uz/en/corporate/certificate_verification?qr_code=%s"
	ruBaseURL = "https://hsegroup.uz/ru/kurumsal/certificate_verification?qr_code=%s"

	dateTimeFormat = "15:04, 02.01.2006"
	location       = "Asia/Samarkand"
)

var (
	UniqueReferenceMustPresentErr = errors.New("unique reference must present")
)

type Record struct {
	FullName           string
	CompanyName        string
	EducationDateStart string
	EducationDateEnd   string
	EducationName      string
	EducationHours     string
	UniqueReference    string
	Language           string
}

// Assigns unique nano id reference code to object.
func (r *Record) AssignUniqueReference() {
	code, _ := gonanoid.Generate(charset, length)

	r.UniqueReference = code
}

// Formats for Google Sheets.
func (r Record) FormatForSheets() []interface{} {
	l, _ := time.LoadLocation(location)

	return []interface {
	}{
		r.FullName,
		r.CompanyName,
		r.EducationName,
		r.EducationHours,
		r.EducationDateStart,
		r.EducationDateEnd,
		r.UniqueReference,
		time.Now().In(l).Format(dateTimeFormat),
	}
}

// Creates QR code for record and returns it as bytes.
func (r Record) GenerateQRCode() ([]byte, error) {
	if r.UniqueReference == "" {
		return nil, UniqueReferenceMustPresentErr
	}

	var baseText string
	switch r.Language {
	case "TR":
		baseText = trBaseURL
	case "EN":
		baseText = enBaseURL
	case "RU":
		baseText = ruBaseURL
	}

	return qrcode.Encode(fmt.Sprintf(baseText, r.UniqueReference), qrcode.Medium, 100)
}
