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
