package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type commandLine struct {
	input textinput.Model
}

func newCommandLine() *commandLine {
	var c commandLine
	c.input = textinput.New()
	c.input.Prompt = ""

	return &c
}

func (l *commandLine) focus() tea.Cmd {
	return l.input.Focus()
}

func (l *commandLine) update(msg tea.Msg) tea.Cmd {
	input, cmd := l.input.Update(msg)
	l.input = input
	return cmd
}

func (l *commandLine) view() string {
	return l.input.View()
}

func (l *commandLine) show() {
	l.input.SetValue("")
	l.input.Prompt = ":"
	l.input.SetCursorMode(textinput.CursorBlink)
}

func (l *commandLine) hide() {
	l.input.Blur()
	l.input.SetValue("")
	l.input.Prompt = ""
	l.input.SetCursorMode(textinput.CursorHide)
}

func (l *commandLine) prompt(command Command) tea.Cmd {
	l.show()
	l.input.Prompt = command.prompt()
	return l.focus()
}
