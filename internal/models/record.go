package models

import (
	"errors"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/skip2/go-qrcode"
	"time"
)

const (
	charset = "123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-"
	length  = 17

	baseURL        = "https://hsegroup.uz/kurumsal/certificate_verification?qr_code="
	dateTimeFormat = "15:04, 02.01.2006"
	location       = "Asia/Samarkand"
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
	Language        string
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
		r.FirstName,
		r.LastName,
		r.CompanyName,
		r.JobName,
		r.EducationName,
		r.EducationHours,
		r.EducationDate,
		r.UniqueReference,
		time.Now().In(l).Format(dateTimeFormat),
	}
}

// Creates QR code for record and returns it as bytes.
func (r Record) GenerateQRCode() ([]byte, error) {
	if r.UniqueReference == "" {
		return nil, UniqueReferenceMustPresentErr
	}

	return qrcode.Encode(baseURL+r.UniqueReference, qrcode.Medium, 100)
}
