package main

import (
	"testing"
)

func TestWordCounterWrite(t *testing.T) {
	counter := new(WordCounter)

	testString := "Hello World, today is Tuesday"
	counter.Write([]byte(testString))

	want := 5
	if int(*counter) != want {
		t.Errorf("Counter did not get correct count, got %v wanted %v\n",
			*counter, want)
	}
}

func TestLineCounterWrite(t *testing.T) {
	counter := new(LineCounter)

	testString := "Hello World\n, today\n is\n Tuesday"
	counter.Write([]byte(testString))

	want := 4
	if int(*counter) != want {
		t.Errorf("Counter did not get correct count, got %v wanted %v\n",
			*counter, want)
	}
}
