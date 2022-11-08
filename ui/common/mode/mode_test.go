package mode

import (
	"testing"
)

func Test_mode_String(t *testing.T) {
	if Command.String() != "COMMAND" {
		t.Fatal("error")
	}
	if Insert.String() != "INSERT" {
		t.Fatal("error")
	}
}
