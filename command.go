package main

type Command interface {
	// will show commandLine prompt
	prompt() string

	// on command is finish will set
	setValue(val string)

	// Value get set be value
	Value() string
}

type SaveFileCommand struct {
	val string
}

func (s *SaveFileCommand) setValue(val string) {
	s.val = val
}

func (s *SaveFileCommand) Value() string {
	return s.val
}

func (s *SaveFileCommand) prompt() string {
	return "Save File: "
}
