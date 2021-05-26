package sheet_handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigitsadic/hsedocumentgenerator/internal/models"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
)

const (
	documentId   = "1fCSvCMXHDrT99ytqY6xZSgk_Wu__NCTNewzZ8rTvW_c"
	pageId       = "Sertifika Yaratıcı"
	dbPageId     = "Sertifika Veritabanı"
	credFileName = "credentials.json"
	pageRange    = "A:H"
)

type SheetHandler struct {
}

func (s SheetHandler) ReadFromSheets() ([]models.Record, error) {
	srv, err := sheets.NewService(
		context.TODO(),
		option.WithScopes(sheets.SpreadsheetsScope),
		option.WithCredentialsFile(credFileName),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	readRange := fmt.Sprintf("%s!%s", pageId, pageRange)
	resp, err := srv.Spreadsheets.Values.Get(documentId, readRange).Do()
	if err != nil {
		return nil, err
	}

	var results []models.Record

	for _, row := range resp.Values[1:] {
		if len(row) < 7 {
			continue
		}

		firstName, ok1 := row[0].(string)
		lastName, ok2 := row[1].(string)
		company, ok3 := row[2].(string)
		job, ok4 := row[3].(string)
		educationName, ok5 := row[4].(string)
		educationDuration, ok6 := row[5].(string)
		educationDate, ok7 := row[6].(string)
		lang, ok8 := row[7].(string)

		if ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8 {
			r := models.Record{
				FirstName:      firstName,
				LastName:       lastName,
				CompanyName:    company,
				JobName:        job,
				EducationName:  educationName,
				EducationHours: educationDuration,
				EducationDate:  educationDate,
				Language:       lang,
			}
			r.AssignUniqueReference()

			results = append(results, r)
		} else {
			continue
		}
	}

	return results, nil
}

func (s SheetHandler) WriteToSheets(records []models.Record) error {
	if len(records) == 0 {
		return errors.New("cannot write to empty records")
	}

	srv, err := sheets.NewService(
		context.TODO(),
		option.WithCredentialsFile(credFileName),
	)
	if err != nil {
		return fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	res, err := srv.Spreadsheets.Values.Get(documentId, dbPageId+"!A:I").Do()
	if err != nil {
		return err
	}

	var vr sheets.ValueRange

	for _, record := range records {
		vr.Values = append(vr.Values, record.FormatForSheets())
	}

	_, err = srv.Spreadsheets.Values.Append(documentId, fmt.Sprintf("%s!A%d", dbPageId, len(res.Values)+1), &vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)

		return err
	}

	return nil
}
