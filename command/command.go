package command

import "strings"

func ParseCommand(cmd string) Command {
	cmd = strings.TrimPrefix(cmd, ":")
	switch cmd {
	case "w":
		return &SaveFile{}
	case "q":
		return &Quit{}
	default:
		return &Unknown{}
	}
}

type Command interface {
	// Prompt will show commandLine prompt
	Prompt() string

	// SetValue on command is finish will set
	SetValue(val string)

	// Value get set be value
	Value() string
}

// Unknown command, means user input cloud not be parsed.
type Unknown struct {
	val string
}

func (u *Unknown) Prompt() string {
	return "unknown command !!!"
}

func (u *Unknown) SetValue(val string) {
	u.val = val
}

func (u *Unknown) Value() string {
	return u.val
}
