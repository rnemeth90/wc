package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	want := 4
	got := count(b, false, false)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1\n word2\n word3\n word4\n")
	want := 4
	got := count(b, false, true)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\n")
	want := 18
	got := count(b, true, false)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
