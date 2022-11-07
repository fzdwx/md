package main

import (
	"testing"
)

func Test_mode_String(t *testing.T) {
	if modeCommand.String() != "COMMAND" {
		t.Fatal("error")
	}
	if modeInsert.String() != "INSERT" {
		t.Fatal("error")
	}
}
