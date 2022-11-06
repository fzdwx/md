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
		"COMMAND",
		"INSERT",
		"NORMAL",
	}[m]
}

const (
	command mode = iota
	insert
	normal
)

type errMsg error

type model struct {
	config             *mdConfig
	showWelcomeContent bool
	err                errMsg

	mode   mode // current mode
	width  int  // the terminal max width
	height int  // the terminal max height

	writeArea       textarea.Model
	welcomeViewPort viewport.Model
	statusLine      statusLine

	md *markdown // current edit markdown file
}

// initialModel
func initialModel(config *mdConfig) *model {
	m := model{}
	m.config = config
	m.showWelcomeContent = config.showWelcomeContent()
	m.mode = normal
	m.md = defaultMd()
	m.statusLine = statusLine{config: config}
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
		m.refreshStatusLine()
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
					m.refreshStatusLine()
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

	if !m.showWelcomeContent {
		m.refreshStatusLine()
	}

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	if m.showWelcomeContent {
		m.writeArea.View() // in my pc, need call this method(very slow)
		return m.welcomeViewPort.View()
	}

	view := m.writeArea.View()
	return view + "\n" + m.statusLine.view() + "\ncommand line"
}

func (m *model) resize(msg tea.WindowSizeMsg) {
	m.width = msg.Width
	m.height = msg.Height
	if m.showWelcomeContent {
		m.welcomeViewPort.Width = msg.Width
		m.welcomeViewPort.Height = msg.Height
	}

	m.writeArea.SetWidth(msg.Width)
	m.writeArea.SetHeight(msg.Height - 2)
}

func (m *model) refreshStatusLine() {
	area := m.writeArea
	m.statusLine.refresh(m.md, m.width, m.height, m.mode, area.Line(), area.LineInfo().ColumnOffset, area.LineCount())
}
