package handlers

import (
	"bytes"
	"strings"
	"testing"
)

func TestHandler_PrintHelloText(t *testing.T) {
	o := new(bytes.Buffer)
	h := Handler{Output: o}

	h.PrintHelloText()

	if !strings.Contains(o.String(), welcomeText) {
		t.Errorf("expected to see %q but got %q", welcomeText, o.String())
	}
}
