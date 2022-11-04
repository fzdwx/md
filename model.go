package main

import (
	_ "embed"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/bubbletea"
	"log"
	"time"
)

//go:embed Default.md
var welcomeContent string

type errMsg error

type mode int

const (
	welcome mode = iota + 1
	normal
	insert
)

type model struct {
	showWelcomeContent bool

	config          *mdConfig
	mode            mode
	writeArea       textarea.Model
	welcomeViewPort viewport.Model
	err             errMsg
}

// initialModel
func initialModel(config *mdConfig) *model {
	m := model{mode: normal}
	m.config = config
	m.showWelcomeContent = config.showWelcomeContent()
	return &m
}

func (m *model) Init() tea.Cmd {
	if m.showWelcomeContent {
		m.mode = welcome
		m.welcomeViewPort = viewport.New(10, 10)
		m.welcomeViewPort.SetContent(mustToMd(welcomeContent, m.config.style))
		return nil
	}
	m.writeArea = textarea.New()
	return textarea.Blink
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.resize(msg)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

		if m.mode == welcome {
			m.mode = insert
			m.writeArea = textarea.New()
			return m, textarea.Blink
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
				cmd = m.writeArea.Focus()
				cmds = append(cmds, cmd)
			}
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	if m.mode == welcome {
		return m, nil
	}

	m.writeArea, cmd = m.writeArea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Sequentially(cmds...)
}

func (m *model) View() string {
	if m.mode == welcome {
		//	return m.welcomeViewPort.View()
		return "asdasd"
	}

	now := time.Now()
	view := m.writeArea.View()
	log.Printf("run textarea view cost: %s", time.Now().Sub(now).String())
	return view
}

func (m *model) resize(msg tea.WindowSizeMsg) {
	if m.mode == welcome {
		m.welcomeViewPort.Width = msg.Width - 10
		m.welcomeViewPort.Height = msg.Height
		return
	}

	m.writeArea.SetWidth(msg.Width - 10)
	m.writeArea.SetHeight(msg.Height)
}
