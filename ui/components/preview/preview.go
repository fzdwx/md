package preview

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/glamour"
	"github.com/fzdwx/md/config"
)

func ToMd(in string, stylePath string, wordWrap int) (string, error) {
	renderer, err := glamour.NewTermRenderer(
		glamour.WithEmoji(),
		glamour.WithStylePath(stylePath),
		glamour.WithWordWrap(wordWrap),
	)
	if err != nil {
		return "", err
	}

	return renderer.Render(in)
}

func MustRender(in string, stylePath string, wordWrap int) string {
	md, _ := ToMd(in, stylePath, wordWrap)
	return md
}

type View struct {
	viewport   viewport.Model
	config     *config.Context
	rawContent string
}

func (v *View) SetContent(content string) {
	v.rawContent = content
	v.RenderMarkdown()
}

func (v *View) View() string {
	return v.viewport.View()
}

func (v *View) Set(width int, height int) {
	v.viewport.Width = width
	v.viewport.Height = height
}

// RenderMarkdown fix wordWrap
func (v *View) RenderMarkdown() {
	v.viewport.SetContent(MustRender(v.rawContent, v.config.MdStyle, v.viewport.Width))
}

func New(config *config.Context) *View {
	p := &View{}
	p.config = config
	p.viewport = viewport.New(0, 0)
	p.viewport.YPosition = -1
	return p
}
