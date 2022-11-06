package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/md/theme"
)

var DefaultStatusLineStyle = statusLineStyle{
	modeStyle:     lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).Foreground(theme.ColorBlank).Background(theme.ColorA1).Bold(true),
	fileNameStyle: lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).Foreground(theme.ColorBlank2),
}

type statusLine struct {
	config    *mdConfig
	row       int
	col       int
	lineCount int
	mode      mode
	width     int
	height    int
	md        *markdown
}

type statusLineStyle struct {
	modeStyle     lipgloss.Style
	fileNameStyle lipgloss.Style
}

func (l *statusLine) refresh(md *markdown, width int, height int, mode mode, row int, col int, lineCount int) {
	l.md = md
	l.width = width
	l.height = height
	l.mode = mode
	l.row = row + 1
	l.col = col + 1
	l.lineCount = lineCount
}

func (l *statusLine) view() string {
	style := l.config.statusLineStyle

	return fmt.Sprintf("%s %s %s %s %s",
		style.modeStyle.Render(l.mode.String()),
		style.fileNameStyle.Render(l.md.fileName),
		fmt.Sprintf("(%d:%d)", l.row, l.col),
		fmt.Sprintf("(%d:%d)", l.width, l.height),
		fmt.Sprintf("total:%d", l.lineCount),
	)
}
