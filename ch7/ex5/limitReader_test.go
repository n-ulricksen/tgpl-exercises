package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	testBytes := []byte("hello world, today it is hot outside")
	var limit int64 = 15

	stringReader := strings.NewReader("")
	limitReader := LimitReader(stringReader, limit)

	bytesRead, err := limitReader.Read(testBytes)
	if len(testBytes) > int(limit) && err == nil {
		t.Errorf("Read() should return EOF error when reading past limit\n")
	}

	fmt.Printf("LOG: Read %d bytes\n", bytesRead)

	if bytesRead > int(limit) {
		t.Errorf("Read() read %d bytes, limit is %d\n", bytesRead, limit)
	}
}
