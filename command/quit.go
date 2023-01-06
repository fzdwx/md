package command

type Quit struct {
}

func (q *Quit) Prompt() string {
	return ""
}

func (q *Quit) SetValue(val string) {
}

func (q *Quit) Value() string {
	return ""
}
