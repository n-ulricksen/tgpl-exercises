package main

import (
	"bufio"
	"bytes"
)

type WordCounter int

func (w *WordCounter) Write(data []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*w++
	}

	return int(*w), nil
}

type LineCounter int

func (l *LineCounter) Write(data []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*l++
	}

	return int(*l), nil
}
