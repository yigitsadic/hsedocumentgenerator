package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/yigitsadic/hsedocumentgenerator/internal/models"
	"reflect"
	"strings"
	"testing"
)

type mockClient struct {
	Output []models.Record
	Error  error
}

func (m mockClient) ReadFromSheets() ([]models.Record, error) {
	return m.Output, m.Error
}

func TestHandler_PrintHelloText(t *testing.T) {
	o := new(bytes.Buffer)
	h := Handler{Output: o}

	h.PrintHelloText()

	if !strings.Contains(o.String(), welcomeText) {
		t.Errorf("expected to see %q but got %q", welcomeText, o.String())
	}
}

func TestHandler_ReadFromSheets(t *testing.T) {
	t.Run("it should handle successful scenario", func(t *testing.T) {
		mC := new(mockClient)
		o := new(bytes.Buffer)
		h := Handler{Output: o, Client: mC}

		mC.Output = []models.Record{
			{
				FirstName:       "Ali",
				LastName:        "Veli",
				JobName:         "Vinç Operatörü",
				CompanyName:     "Bir Şirket",
				EducationDate:   "07.02.2012",
				EducationName:   "Güvenli Sürüş",
				EducationHours:  "22 saat",
				UniqueReference: "ABCDEF",
			},
		}
		mC.Error = nil

		h.ReadFromSheets()

		if !strings.Contains(o.String(), fmt.Sprintf(recordReadText, len(mC.Output))) {
			t.Errorf("expected output not satisfied")
		}

		if !reflect.DeepEqual(mC.Output, h.ReadRecords) {
			t.Errorf("expected to read records assign")
		}
	})

	t.Run("it should handle failure scenario", func(t *testing.T) {
		mC := new(mockClient)
		o := new(bytes.Buffer)
		h := Handler{Output: o, Client: mC}

		mC.Output = nil
		mC.Error = errors.New("some error occurred")

		err := h.ReadFromSheets()
		if err == nil {
			t.Errorf("expected to see error but got nothing")
		}

		if !strings.Contains(o.String(), cannotReadFromGoogleText) {
			t.Errorf("expected output not satisfied. expected=%q but got=%q", cannotReadFromGoogleText, o.String())
		}
	})
}
