package preview

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/glamour"
	"github.com/fzdwx/md/config"
)

func ToMd(in string, stylePath string) (string, error) {
	return glamour.Render(in, stylePath)
}

func MustRender(in string, stylePath string) string {
	md, _ := ToMd(in, stylePath)
	return md
}

type View struct {
	viewport viewport.Model
	config   *config.Context
}

func (v *View) SetContent(content string) {
	v.viewport.SetContent(MustRender(content, v.config.MdStyle))
}

func (v *View) View() string {
	return v.viewport.View()
}

func (v *View) Set(width int, height int) {
	v.viewport.Width = width
	v.viewport.Height = height
}

func New(config *config.Context) *View {
	p := &View{}
	p.config = config
	p.viewport = viewport.New(0, 0)
	return p
}
