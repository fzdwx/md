package main

import (
	_ "embed"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/md/command"
	"github.com/fzdwx/md/utils"
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
	modeCommand mode = iota
	modeInsert
	modeNormal
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
	m.mode = modeNormal
	m.md = defaultMd()
	m.statusLine = &statusLine{config: config}
	m.commandLine = newCommandLine(config)
	return &m
}

func (m *model) Init() tea.Cmd {
	m.writeArea = textarea.New()
	m.writeArea.KeyMap = m.config.keymap.InsertModeKeyMap

	if m.showWelcomeContent {
		m.welcomeViewPort = viewport.New(0, 0)
		m.md = defaultMd()
		m.welcomeViewPort.SetContent(mustRender(welcomeContent, m.config.mdStyle))
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
	case command.Command:
		m.toNormalMode()
		switch msg.(type) {
		case *command.SaveFile:
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
					return m, m.prompt(&command.SaveFile{}) // todo 使用 modeCommand 模式 获取输入内容
				}
				m.savefile()
				return m, nil
			}
		case key.Matches(msg, DefaultKeyMap.ToCommandMode):
			//if m.mode == modeNormal && !m.showWelcomeContent {
			//	m.toCommandMode()
			//	return m, m.commandLine.focus()
			//}
			// allow in welcome view use command mode
			m.toCommandMode()
			return m, m.commandLine.focus()
		case key.Matches(msg, DefaultKeyMap.ToNormalMode):
			m.toNormalMode()
		case key.Matches(msg, DefaultKeyMap.ToInsertMode):
			if m.mode == modeNormal {
				m.toInsertMode()
				return m, nil
			}
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	switch m.mode {
	case modeInsert:
		m.writeArea, cmd = m.writeArea.Update(msg)
		cmds = append(cmds, cmd)
	case modeCommand:
		cmd = m.commandLine.update(msg)
		cmds = append(cmds, cmd)
	}

	if !m.showWelcomeContent {
		m.refreshStatusLine()
	}

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	buffer := utils.NewStrBuffer()
	if m.showWelcomeContent {
		m.writeArea.View() // in my pc, need call this method(very slow)
		buffer.Write(m.welcomeViewPort.View())
	} else {
		buffer.Write(m.writeArea.View()).NewLine().Write(m.statusLine.view())
	}

	return buffer.NewLine().Write(m.commandLine.view()).String()
}

func (m *model) resize(msg tea.WindowSizeMsg) {
	m.width = msg.Width
	m.height = msg.Height
	if m.showWelcomeContent {
		m.welcomeViewPort.Width = msg.Width
		m.welcomeViewPort.Height = msg.Height - 2
	}

	m.writeArea.SetWidth(msg.Width)
	m.writeArea.SetHeight(msg.Height - 2)
}

func (m *model) toCommandMode() {
	m.mode = modeCommand
	m.commandLine.show()
	m.refreshStatusLine()
}

func (m *model) toInsertMode() {
	m.mode = modeInsert
	if m.showWelcomeContent {
		m.showWelcomeContent = false
	}

	m.writeArea.Focus()
	m.refreshStatusLine()
}

func (m *model) toNormalMode() {
	m.mode = modeNormal
	m.writeArea.Blur()
	m.commandLine.hide()
}

func (m *model) prompt(cmd command.Command) tea.Cmd {
	m.mode = modeCommand
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
