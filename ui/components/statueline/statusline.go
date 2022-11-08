package statueline

import (
	"fmt"
	"github.com/fzdwx/md/config"
	"github.com/fzdwx/md/icon"
	"github.com/fzdwx/md/ui/common/doc"
	"github.com/fzdwx/md/ui/common/mode"
)

func New(config *config.Context) *Bar {
	return &Bar{config: config}
}

type Bar struct {
	config    *config.Context
	row       int
	col       int
	lineCount int
	mode      mode.Mode
	width     int
	height    int
	md        *doc.Markdown
}

func (l *Bar) Refresh(md *doc.Markdown, width int, height int, mode mode.Mode, row int, col int, lineCount int) {
	l.md = md
	l.width = width
	l.height = height
	l.mode = mode
	l.row = row + 1
	l.col = col + 1
	l.lineCount = lineCount
}

func (l *Bar) View() string {
	style := l.config.Style.StatusLineStyle

	return fmt.Sprintf("%s %s %s %s %s %s",
		style.ModeStyle.Render(l.mode.String()),
		style.FileNameStyle.Render(l.md.FileName),
		fmt.Sprintf("(%d:%d)", l.row, l.col),
		fmt.Sprintf("(%d:%d)", l.width, l.height),
		fmt.Sprintf("total:%d", l.lineCount),
		icon.Markdown,
	)
}
