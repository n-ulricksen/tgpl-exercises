package main

import (
	"strings"
	"testing"
)

func TestNewReader(t *testing.T) {
	r := NewReader()

	testWords := []string{"hello ", "world ", "goodbye ", "world "}
	for _, word := range testWords {
		r.Read([]byte(word))
	}

	joined := strings.Join(testWords, "")
	if r.String() != joined {
		t.Errorf("NewReader().Read() error: got '%v', wanted '%v'\n", *r, joined)
	}
}
