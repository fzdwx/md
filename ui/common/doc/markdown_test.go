package doc

import (
	"testing"
)

func Test_FilePathToMd(t *testing.T) {
	file := "Default.md"
	_, err := FilePathToMd(file)
	if err != nil {
		t.Fatal(file, " LoadBody to Markdown fail:", err)
	}
}

func Test_FilePathToMd2(t *testing.T) {
	file := "Defaultasd.md"
	_, err := FilePathToMd(file)
	if err == nil {
		t.Fatal(file, " LoadBody to Markdown fail:", err)
	}
}

func TestNoname(t *testing.T) {
	md := DefaultMd()

	if !md.NoName() {
		t.Fatal("NoName is not expected")
	}
}
