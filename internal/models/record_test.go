package models

import (
	"testing"
)

func TestRecord_AssignUniqueReference(t *testing.T) {
	r := Record{}

	if r.UniqueReference != "" {
		t.Errorf("expected to empty unique reference but got=%q", r.UniqueReference)
	}

	r.AssignUniqueReference()

	if len(r.UniqueReference) != length {
		t.Errorf("unexpected to see %d char unique reference expected was=%d", len(r.UniqueReference), length)
	}
}

func TestRecord_GenerateQRCode(t *testing.T) {
	t.Run("it should give error if reference code is blank", func(t *testing.T) {
		r := Record{UniqueReference: ""}

		_, err := r.GenerateQRCode()

		if err != UniqueReferenceMustPresentErr {
			t.Errorf("expected to get %s error but got %s", UniqueReferenceMustPresentErr, err)
		}
	})

	t.Run("it should generate qr code", func(t *testing.T) {
		r := Record{UniqueReference: "ABCDEF"}

		_, err := r.GenerateQRCode()

		if err != nil {
			t.Errorf("unexpected to get an error but got=%s", err)
		}
	})
}

func TestRecord_FormatForSheets(t *testing.T) {
	r := Record{
		FullName:        "Yigit Sadicis",
		CompanyName:     "Medya ve Toplum",
		EducationDate:   "20.07.2022",
		EducationName:   "Reklam ve YouTube",
		EducationHours:  "96 saat",
		UniqueReference: "ylq123-2eiQ",
	}

	got := r.FormatForSheets()

	if v, ok := got[0].(string); !ok || v != r.FullName {
		t.Errorf("expected value not satisfied")
	}

	if v, ok := got[1].(string); !ok || v != r.CompanyName {
		t.Errorf("expected value not satisfied")
	}

	if v, ok := got[2].(string); !ok || v != r.EducationName {
		t.Errorf("expected value not satisfied")
	}

	if v, ok := got[3].(string); !ok || v != r.EducationHours {
		t.Errorf("expected value not satisfied")
	}

	if v, ok := got[4].(string); !ok || v != r.EducationDate {
		t.Errorf("expected value not satisfied")
	}

	if v, ok := got[5].(string); !ok || v != r.UniqueReference {
		t.Errorf("expected value not satisfied")
	}
}
