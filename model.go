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
	commandLine     *commandLine
	statusLine      *statusLine

	md *markdown // current edit markdown file
}

// initialModel
func initialModel(config *mdConfig) *model {
	m := model{}
	m.config = config
	m.showWelcomeContent = config.showWelcomeContent()
	m.mode = normal
	m.md = defaultMd()
	m.statusLine = &statusLine{config: config}
	m.commandLine = newCommandLine(config)
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
	case Command:
		m.toNormalMode()
		switch msg.(type) {
		case *SaveFileCommand:
			m.md.fileName = msg.Value()
			m.savefile()
			m.refreshStatusLine()
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.resize(msg)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.config.keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.config.keymap.SaveFile):
			if !m.showWelcomeContent {
				if m.md.noName() {
					return m, m.prompt(&SaveFileCommand{}) // todo 使用 command 模式 获取输入内容
				}
				m.savefile()
				return m, nil
			}
		case key.Matches(msg, DefaultKeyMap.ToCommandMode):
			if m.mode == normal && !m.showWelcomeContent {
				m.toCommandMode()
				return m, m.commandLine.focus()
			}
		case key.Matches(msg, DefaultKeyMap.ToNormalMode):
			m.toNormalMode()
		case key.Matches(msg, DefaultKeyMap.ToInsertMode):
			if m.mode == normal {
				m.toInsertMode()
				return m, nil
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
		cmd = m.commandLine.update(msg)
		cmds = append(cmds, cmd)
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
	return view + "\n" + m.statusLine.view() + "\n" + m.commandLine.view()
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

func (m *model) toCommandMode() {
	m.mode = command
	m.commandLine.show()
	m.refreshStatusLine()
}

func (m *model) toInsertMode() {
	m.mode = insert
	if m.showWelcomeContent {
		m.showWelcomeContent = false
	}

	m.writeArea.Focus()
	m.refreshStatusLine()
}

func (m *model) toNormalMode() {
	m.mode = normal
	m.writeArea.Blur()
	m.commandLine.hide()
}

func (m *model) prompt(cmd Command) tea.Cmd {
	m.mode = command
	m.writeArea.Blur()
	m.refreshStatusLine()
	return m.commandLine.prompt(cmd)
}

func (m *model) hideCommandLine() {
	m.commandLine.hide()
}

func (m *model) refreshStatusLine() {
	area := m.writeArea
	m.statusLine.refresh(m.md, m.width, m.height, m.mode, area.Line(), area.LineInfo().ColumnOffset, area.LineCount())
}

func (m *model) savefile() {
	m.md.body = m.writeArea.Value()
	m.err = m.md.save()
}
