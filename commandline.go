package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type commandLine struct {
	input  textinput.Model
	config *mdConfig

	cmd Command
}

func newCommandLine(config *mdConfig) *commandLine {
	var c commandLine
	c.input = textinput.New()
	c.input.Prompt = ""
	c.config = config

	return &c
}

func (l *commandLine) focus() tea.Cmd {
	return l.input.Focus()
}

func (l *commandLine) update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		// prompt handle finish, tell main model, do next action.
		case key.Matches(msg, l.config.keymap.commandLineKeyMap.Cr):
			l.cmd.setValue(l.input.Value())
			return func() tea.Msg {
				return l.cmd
			}
		}
	}

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
	l.cmd = nil
	l.input.SetCursorMode(textinput.CursorBlink)
}

func (l *commandLine) hide() {
	l.input.Blur()
	l.input.SetValue("")
	l.input.Prompt = ""
	l.cmd = nil
	l.input.SetCursorMode(textinput.CursorHide)
}

func (l *commandLine) prompt(command Command) tea.Cmd {
	l.show()
	l.cmd = command
	l.input.Prompt = command.prompt()
	return l.focus()
}
