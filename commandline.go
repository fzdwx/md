package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/md/command"
)

type commandLine struct {
	input  textinput.Model
	config *mdConfig

	cmd command.Command
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
		case key.Matches(msg, l.config.keymap.CommandLineKeyMap.Cr):
			// todo modeCommand dispatch
			cmd := l.dispatch()
			return func() tea.Msg {
				return cmd
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

// prompt focus someone Command,
// get teh user input and send it to the main ui to execute it.
func (l *commandLine) prompt(command command.Command) tea.Cmd {
	l.show()
	l.cmd = command
	l.input.Prompt = command.Prompt()
	return l.focus()
}

// dispatch user press CommandLineKeyMap.Cr,
// means that the user has confirmed the input,
// so we have to dispatch to the specific Command.
func (l *commandLine) dispatch() command.Command {
	if l.cmd != nil {
		l.cmd.SetValue(l.input.Value())
		return l.cmd
	}

	// todo dispatch
	return &command.Unknown{}
}
