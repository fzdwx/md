package main

import (
	"testing"
)

func Test_mode_String(t *testing.T) {
	if command.String() != "COMMAND" {
		t.Fatal("error")
	}
	if insert.String() != "INSERT" {
		t.Fatal("error")
	}
}
