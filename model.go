package main

import (
	_ "embed"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/bubbletea"
)

//go:embed Default.md
var welcomeContent string

type mode int

func (m mode) String() string {
	return []string{
		"command",
		"insert",
	}[m]
}

const (
	command mode = iota
	insert
	normal
)

type errMsg error

type model struct {
	showWelcomeContent bool
	mode               mode

	config          *mdConfig
	writeArea       textarea.Model
	welcomeViewPort viewport.Model
	err             errMsg

	md *markdown
}

// initialModel
func initialModel(config *mdConfig) *model {
	m := model{}
	m.config = config
	m.showWelcomeContent = config.showWelcomeContent()
	m.mode = normal
	return &m
}

func (m *model) Init() tea.Cmd {
	m.writeArea = textarea.New()
	m.writeArea.KeyMap = m.config.keymap.insertModeKeyMap

	if m.showWelcomeContent {
		m.welcomeViewPort = viewport.New(0, 0)
		md := mustRender(welcomeContent, m.config.mdStyle)
		m.welcomeViewPort.SetContent(md)
		return nil
	} else {
		m.md, m.err = filePathToMd(m.config.filePath)
		m.writeArea.SetValue(m.md.body)
	}

	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.resize(msg)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.config.keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap.ToCommandMode):
			if m.mode == normal {
				m.mode = command
			}
		case key.Matches(msg, DefaultKeyMap.ToNormalMode):
			m.mode = normal
			if m.writeArea.Focused() {
				m.writeArea.Blur()
			}
		case key.Matches(msg, DefaultKeyMap.ToInsertMode):
			if m.mode == normal {
				m.mode = insert
				if m.showWelcomeContent {
					m.showWelcomeContent = false
				}

				if !m.writeArea.Focused() {
					m.writeArea.Focus()
					return m, nil
				}
			}
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	switch m.mode {
	case insert:
		m.writeArea, cmd = m.writeArea.Update(msg)
		cmds = append(cmds, cmd)
	case command:
		// todo handle command
	}

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	if m.showWelcomeContent {
		return m.welcomeViewPort.View()
	}

	view := m.writeArea.View()
	return view
}

func (m *model) resize(msg tea.WindowSizeMsg) {
	if m.showWelcomeContent {
		m.welcomeViewPort.Width = msg.Width
		m.welcomeViewPort.Height = msg.Height
	}

	m.writeArea.SetWidth(msg.Width)
	m.writeArea.SetHeight(msg.Height)
}
