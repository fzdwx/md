package command

type SaveFile struct {
	val string
}

func (s *SaveFile) SetValue(val string) {
	s.val = val
}

func (s *SaveFile) Value() string {
	return s.val
}

func (s *SaveFile) Prompt() string {
	return "Save File: "
}
