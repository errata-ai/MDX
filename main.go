package main

import (
	"errors"
	"os"
	"path/filepath"
)

func main() {
	expected := filepath.Join("testdata", "in.md")
	observed := filepath.Join("testdata", "out.md")

	// Read the expected and observed files
	expectedContent, err := os.ReadFile(expected)
	if err != nil {
		panic(err)
	}

	observedContent, err := os.ReadFile(observed)
	if err != nil {
		panic(err)
	}

	// Compare the contents of the files
	if string(expectedContent) != string(observedContent) {
		panic(errors.New("expected and observed files do not match"))
	}
}
