package main

import (
	"bytes"
	"os"
	"testing"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name       string
		words      string
		countLines bool
		countBytes bool
		expected   int
	}{
		{name: "test count words", words: "this is a test", countLines: false, countBytes: false, expected: 4},
		{name: "test count lines", words: "this is a test\nand another", countLines: true, countBytes: false, expected: 2},
		{name: "test count bytes", words: "this is another test", countLines: false, countBytes: true, expected: 20},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			bufferString := bytes.NewBufferString(test.words)

			want := test.expected
			got, _ := count(bufferString, test.countLines, test.countBytes)

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func TestReadFiles(t *testing.T) {
	file := "./testFiles/README.md"

	testCases := []struct {
		name       string
		countLines bool
		countBytes bool
		expected   int
	}{
		{name: "test count bytes", countLines: false, countBytes: false, expected: 8},
		{name: "test count bytes", countLines: false, countBytes: true, expected: 34},
		{name: "test count bytes", countLines: true, countBytes: false, expected: 6},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(file)
			if err != nil {
				t.Fatalf("cannot open test file %s", file)
			}

			want := test.expected
			got, _ := count(f, test.countLines, test.countBytes)

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
