package main

import (
	"log"
	"testing"
)

func TestRunCommand(t *testing.T) {
	actual, err := runCommand("abc", []string{"cat", "-"})
	if err != nil {
		log.Fatal(err)
	}
	expected := "abc"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestRunLikePipe(t *testing.T) {
	raw := `abc
	def
	ghi`

	actual, err := runLikePipe([]string{raw, "rev", "head -1"})
	if err != nil {
		log.Fatal(err)
	}
	expected := "cba"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
