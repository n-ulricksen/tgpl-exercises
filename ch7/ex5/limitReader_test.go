package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	testString := "hello world, today it is hot outside"
	var limit int64 = 15

	stringReader := strings.NewReader(testString)
	limitReader := LimitReader(stringReader, limit)

	bytesRead := make([]byte, 50)
	nBytesRead, err := limitReader.Read(bytesRead)
	if (nBytesRead > int(limit)) && err == nil {
		t.Errorf("Read() should return EOF error when reading past limit\n")
	}

	fmt.Printf("LOG: Read %d bytes\n", nBytesRead)
	fmt.Printf("LOG: %s\n", bytesRead)
}
