package mode

type Mode int

func (m Mode) String() string {
	return []string{
		"COMMAND",
		"INSERT",
		"NORMAL",
	}[m]
}

const (
	Command Mode = iota
	Insert
	Normal
)
