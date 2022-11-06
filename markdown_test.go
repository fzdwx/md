package main

import (
	"testing"
)

func Test_filePathToMd(t *testing.T) {
	file := "Default.md"
	_, err := filePathToMd(file)
	if err != nil {
		t.Fatal(file, " loadBody to markdown fail:", err)
	}
}

func Test_filePathToMd2(t *testing.T) {
	file := "Defaultasd.md"
	_, err := filePathToMd(file)
	if err == nil {
		t.Fatal(file, " loadBody to markdown fail:", err)
	}
}

func TestNoname(t *testing.T) {
	md := defaultMd()

	if !md.noName() {
		t.Fatal("noName is not expected")
	}
}
