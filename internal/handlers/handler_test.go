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

func TestHandler_WriteToConsole(t *testing.T) {
	t.Run("it should write hello text", func(t *testing.T) {
		o := new(bytes.Buffer)
		h := Handler{Output: o}

		h.Write(welcomeText)

		if !strings.Contains(o.String(), welcomeText) {
			t.Errorf("expected to see %q but got %q", welcomeText, o.String())
		}
	})

	t.Run("it should write with parameters", func(t *testing.T) {
		o := new(bytes.Buffer)
		h := Handler{Output: o}

		h.Write(recordReadText, 1)

		expected := fmt.Sprintf(recordReadText, 1)

		if !strings.Contains(o.String(), expected) {
			t.Errorf("expected output not satisfied. expected=%q but got=%q", expected, o.String())
		}
	})
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

func TestHandler_StoreOutputPath(t *testing.T) {
	filePath := "myoutput.zip"

	output := new(bytes.Buffer)
	input := new(bytes.Buffer)

	h := NewHandler(input, output, nil)

	input.WriteString(filePath)
	h.StoreOutputPath()

	if !strings.Contains(output.String(), outputZIPText) {
		t.Errorf("expected output not seen")
	}

	if h.ZipOutputPath != filePath {
		t.Errorf("expected zip file path was=%q but got=%q", filePath, h.ZipOutputPath)
	}
}
