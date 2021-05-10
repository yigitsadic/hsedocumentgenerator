package models

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"log"
)

const (
	charset = "123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-"
	length  = 17
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

func (r *Record) AssignUniqueReference() {
	code, err := gonanoid.Generate(charset, length)
	if err != nil {
		log.Fatalln("Error occurred while assigning unique reference")

		return
	}

	r.UniqueReference = code
}
