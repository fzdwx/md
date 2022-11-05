package main

import (
	"testing"
)

func Test_filePathToMd(t *testing.T) {
	file := "Default.md"
	_, err := filePathToMd(file)
	if err != nil {
		t.Fatal(file, " load to markdown fail:", err)
	}
}

func Test_filePathToMd2(t *testing.T) {
	file := "Defaultasd.md"
	_, err := filePathToMd(file)
	if err == nil {
		t.Fatal(file, " load to markdown fail:", err)
	}
}
