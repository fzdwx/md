package style

import (
	"github.com/charmbracelet/lipgloss"
)

type Style struct {
	StatusLineStyle StatusLineStyle
}

type StatusLineStyle struct {
	ModeStyle     lipgloss.Style
	FileNameStyle lipgloss.Style
}

var DefaultStyle = Style{StatusLineStyle: StatusLineStyle{
	ModeStyle:     lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).Foreground(ColorBlank).Background(ColorA1).Bold(true),
	FileNameStyle: lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).Foreground(ColorBlank2),
}}
