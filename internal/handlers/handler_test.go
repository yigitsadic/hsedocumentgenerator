package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/thecodingmachine/gotenberg-go-client/v7"
	"github.com/yigitsadic/hsedocumentgenerator/internal/models"
	"reflect"
	"strings"
	"testing"
)

type mockClient struct {
	Output []models.Record
	Error  error
}

func (m mockClient) WriteToSheets(records []models.Record) error {
	return nil
}

func (m mockClient) ReadFromSheets() ([]models.Record, error) {
	return m.Output, m.Error
}

type mockPDFGenerator struct {
	BuildRequestError error
	BuildError        error
}

func (m mockPDFGenerator) Ping() error {
	return nil
}

func (m mockPDFGenerator) Build(req *gotenberg.HTMLRequest) ([]byte, error) {
	if m.BuildError != nil {
		return nil, m.BuildError
	}

	return []byte("a"), nil
}

func (m mockPDFGenerator) BuildRequest(r models.Record) (*gotenberg.HTMLRequest, error) {
	if m.BuildRequestError != nil {
		return nil, m.BuildRequestError
	}

	return nil, nil
}

type mockFileZipper struct {
	Error error
}

func (m mockFileZipper) WriteAsZip(fileName string, files []models.ReadFile) error {
	if m.Error != nil {
		return m.Error
	}

	return nil
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

	h := NewHandler(input, output, nil, nil, nil)

	input.WriteString(filePath)
	h.StoreOutputPath()

	if !strings.Contains(output.String(), outputZIPText) {
		t.Errorf("expected output not seen")
	}

	if h.ZipOutputPath != filePath {
		t.Errorf("expected zip file path was=%q but got=%q", filePath, h.ZipOutputPath)
	}
}

func TestHandler_GeneratePDF(t *testing.T) {
	t.Run("it should handle build request error gracefully", func(t *testing.T) {
		expectedError := errors.New("basic error")

		h := Handler{}
		h.PDFGenerator = mockPDFGenerator{BuildRequestError: expectedError}
		h.Files = []models.ReadFile{}
		r := models.Record{UniqueReference: "LOREM"}

		err := h.GeneratePDF(r)

		if err != expectedError {
			t.Errorf("expected error was %s but got=%s", expectedError, err)
		}

		if len(h.Files) != 0 {
			t.Errorf("unexpected to file written into files")
		}
	})

	t.Run("it should handle build error gracefully", func(t *testing.T) {
		expectedError := errors.New("basic error")

		h := Handler{}
		h.PDFGenerator = mockPDFGenerator{BuildError: expectedError}
		h.Files = []models.ReadFile{}
		r := models.Record{UniqueReference: "LOREM"}

		err := h.GeneratePDF(r)

		if err != expectedError {
			t.Errorf("expected error was %s but got=%s", expectedError, err)
		}

		if len(h.Files) != 0 {
			t.Errorf("unexpected to file written into files")
		}
	})

	t.Run("it should write to files successfully", func(t *testing.T) {
		o := new(bytes.Buffer)

		h := Handler{Output: o}
		h.PDFGenerator = mockPDFGenerator{}
		h.Files = []models.ReadFile{}
		r := models.Record{UniqueReference: "LOREM", FirstName: "Lorem", LastName: "Ipsum"}

		err := h.GeneratePDF(r)

		if err != nil {
			t.Errorf("unexpected to get an error but got=%s", err)
		}

		if len(h.Files) != 1 {
			t.Errorf("expected to file written into files")
		}

		expectedFileName := fmt.Sprintf("%s.pdf", r.UniqueReference)

		if h.Files[0].FileName != expectedFileName {
			t.Errorf("expected file name was=%s but got=%s", expectedFileName, h.Files[0].FileName)
		}

		expectedText := fmt.Sprintf(pdfGeneratedText, r.UniqueReference, r.FirstName, r.LastName)

		if !strings.Contains(o.String(), expectedText) {
			t.Errorf("expected output not satisfied. expected=%q but got=%q", expectedText, o.String())
		}
	})
}

func TestHandler_WriteFilesToZip(t *testing.T) {
	t.Run("it should handle error gracefully", func(t *testing.T) {
		o := new(bytes.Buffer)
		expectedError := errors.New("hello I am an error")

		h := Handler{Output: o, ZipOutputPath: "example.zip", ZipWriter: mockFileZipper{Error: expectedError}}

		err := h.WriteFilesToZip()
		if err != expectedError {
			t.Errorf("expected error not satisfied. expected=%s but got=%s", expectedError, err)
		}

		if o.String() != "" {
			t.Errorf("unexpected to see output but got=%q", o.String())
		}
	})

	t.Run("it should work as intentend", func(t *testing.T) {
		o := new(bytes.Buffer)

		h := Handler{Output: o, ZipOutputPath: "example.zip", ZipWriter: mockFileZipper{Error: nil}}

		err := h.WriteFilesToZip()
		if err != nil {
			t.Errorf("unexpected to see an error but got=%s", err)
		}

		expectedOutput := fmt.Sprintf(zipFileCreatedText, "example.zip")

		if !strings.Contains(o.String(), expectedOutput) {
			t.Errorf("expected output not satisfied")
		}
	})
}

func TestHandler_Do(t *testing.T) {
	t.Run("it should work as intented", func(t *testing.T) {
		expectedOutput := `🚀	Google Sheets üzerinden okuma başlatıldı.
📗	Google Sheets üzerinden 2 kayıt okundu.
🤔	Oluşturulan PDFleri içeren ZIP dosyası nereye depolanacak?:	⏳	PDF belge üretme işlemi başlandı...
👍	[abc.pdf]	Lorem	Ipsum	için PDF belgesi üretildi.
👍	[def.pdf]	Ali	Veli	için PDF belgesi üretildi.
✅	PDF belgeleri "example.csv" olarak sıkıştırıldı ve okunan kayıtlar Google Sheets içine eklendi.
💫	İşlem tamamlandı. İyi günler!
`

		o := new(bytes.Buffer)
		i := new(bytes.Buffer)

		records := []models.Record{
			{
				FirstName:       "Lorem",
				LastName:        "Ipsum",
				UniqueReference: "abc",
			},
			{
				FirstName:       "Ali",
				LastName:        "Veli",
				UniqueReference: "def",
			},
		}

		h := NewHandler(i, o, mockClient{Output: records}, mockPDFGenerator{}, mockFileZipper{Error: nil})

		i.WriteString("example.csv\n")

		h.Do()

		if !strings.Contains(o.String(), expectedOutput) {
			t.Errorf("expected output not satisfied")
		}
	})

	t.Run("it should handle with no record found", func(t *testing.T) {
		expectedOutput := `🚀	Google Sheets üzerinden okuma başlatıldı.
📗	Google Sheets üzerinden 0 kayıt okundu.
🥺	Google Sheets üzerinde kayıt bulunamadı. Yapacak bir şey yok.
`

		o := new(bytes.Buffer)
		i := new(bytes.Buffer)

		h := NewHandler(i, o, mockClient{Output: nil}, mockPDFGenerator{}, nil)

		i.WriteString("example.csv\n")

		h.Do()

		if !strings.Contains(o.String(), expectedOutput) {
			t.Errorf("expected output not satisfied")
		}
	})

	t.Run("it should handle when error occurred", func(t *testing.T) {
		expectedOutput := `🚀	Google Sheets üzerinden okuma başlatıldı.
📗	Google Sheets üzerinden 2 kayıt okundu.
🤔	Oluşturulan PDFleri içeren ZIP dosyası nereye depolanacak?:	⏳	PDF belge üretme işlemi başlandı...
😥	[abc.pdf] Lorem Ipsum için beklenmedik bir hata oluştu.
😥	[def.pdf] Ali Veli için beklenmedik bir hata oluştu.
🙈	Sıkıştırılacak PDF bulunamadı.`

		expectedError := errors.New("hello expected error here")
		o := new(bytes.Buffer)
		i := new(bytes.Buffer)
		records := []models.Record{
			{
				FirstName:       "Lorem",
				LastName:        "Ipsum",
				UniqueReference: "abc",
			},
			{
				FirstName:       "Ali",
				LastName:        "Veli",
				UniqueReference: "def",
			},
		}

		h := NewHandler(i, o, mockClient{Output: records}, mockPDFGenerator{BuildError: expectedError}, mockFileZipper{})
		i.WriteString("example.csv\n")
		h.Do()

		if !strings.Contains(o.String(), expectedOutput) {
			t.Errorf("expected output not satisfied")
		}
	})
}
