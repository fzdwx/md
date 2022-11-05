package main

import (
	"testing"
)

func Test_mode_String(t *testing.T) {
	if command.String() != "command" {
		t.Fatal("error")
	}
	if insert.String() != "insert" {
		t.Fatal("error")
	}
}
