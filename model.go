package main

import (
	_ "embed"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/bubbletea"
)

//go:embed Default.md
var welcomeContent string

type errMsg error

type model struct {
	showWelcomeContent bool

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
	return &m
}

func (m *model) Init() tea.Cmd {
	m.writeArea = textarea.New()

	if m.showWelcomeContent {
		m.welcomeViewPort = viewport.New(0, 0)
		md := mustRender(welcomeContent, m.config.style)
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
		m.showWelcomeContent = false

		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

		switch msg.Type {
		case tea.KeyEsc:
			if m.writeArea.Focused() {
				m.writeArea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.writeArea.Focused() {
				m.writeArea.Focus()
			}
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.writeArea, cmd = m.writeArea.Update(msg)
	cmds = append(cmds, cmd)
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
