package main

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var buf bytes.Buffer
	writer, count := CountingWriter(&buf)

	testString := "hello world"
	l := int64(len(testString))

	testRuns := 5
	for i := 1; i <= testRuns; i++ {
		writer.Write([]byte(testString))
		if *count != l*int64(i) {
			t.Errorf("Count should update after Write() call, count = %d\n",
				*count)
		}
	}
}
