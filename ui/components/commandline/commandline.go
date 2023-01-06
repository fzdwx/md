package commandline

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/md/command"
	"github.com/fzdwx/md/config"
)

type Bar struct {
	input  textinput.Model
	config *config.Context

	cmd command.Command
}

func New(config *config.Context) *Bar {
	var c Bar
	c.input = textinput.New()
	c.input.Prompt = ""
	c.config = config

	return &c
}

func (l *Bar) Focus() tea.Cmd {
	return l.input.Focus()
}

func (l *Bar) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		// Prompt handle finish, tell main model, do next action.
		case key.Matches(msg, l.config.Keymap.CommandLineKeyMap.Cr):
			// todo modeCommand Dispatch
			cmd := l.Dispatch()
			return func() tea.Msg {
				return cmd
			}
		}
	}

	input, cmd := l.input.Update(msg)
	l.input = input
	return cmd
}

func (l *Bar) View() string {
	return l.input.View()
}

func (l *Bar) Show() {
	l.input.SetValue("")
	l.input.Prompt = ":"
	l.cmd = nil
	l.input.Cursor.SetMode(cursor.CursorBlink)
}

func (l *Bar) Hide() {
	l.input.Blur()
	l.input.SetValue("")
	l.input.Prompt = ""
	l.cmd = nil
	l.input.Cursor.SetMode(cursor.CursorHide)
}

// Prompt Focus someone Command,
// get teh user input and send it to the main ui to execute it.
func (l *Bar) Prompt(command command.Command) tea.Cmd {
	l.Show()
	l.cmd = command
	l.input.Prompt = command.Prompt()
	return l.Focus()
}

// Dispatch user press CommandLineKeyMap.Cr,
// means that the user has confirmed the input,
// so we have to Dispatch to the specific Command.
func (l *Bar) Dispatch() command.Command {
	value := l.input.Value()
	if l.cmd != nil {
		l.cmd.SetValue(value)
		return l.cmd
	}

	return command.ParseCommand(value)
}
