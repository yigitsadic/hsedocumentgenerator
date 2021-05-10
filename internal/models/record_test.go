package models

import "testing"

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
