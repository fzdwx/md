package main

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/glamour"
)

func toMd(in string, stylePath string) (string, error) {
	return glamour.Render(in, stylePath)
}

func mustRender(in string, stylePath string) string {
	md, _ := toMd(in, stylePath)
	return md
}

type previewView struct {
	viewport viewport.Model
	config   *mdConfig
}

func (v *previewView) SetContent(content string) {
	v.viewport.SetContent(mustRender(content, v.config.mdStyle))
}

func (v *previewView) View() string {
	return v.viewport.View()
}

func (v *previewView) Set(width int, height int) {
	v.viewport.Width = width
	v.viewport.Height = height
}

func newPreviewView(config *mdConfig) *previewView {
	p := &previewView{}
	p.config = config
	p.viewport = viewport.New(0, 0)
	return p
}
