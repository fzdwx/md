package ui

import (
	_ "embed"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/md/command"
	"github.com/fzdwx/md/config"
	"github.com/fzdwx/md/ui/common/doc"
	"github.com/fzdwx/md/ui/common/mode"
	"github.com/fzdwx/md/ui/components/commandline"
	"github.com/fzdwx/md/ui/components/preview"
	"github.com/fzdwx/md/ui/components/statueline"
	"github.com/fzdwx/md/utils"
)

//go:embed Default.md
var welcomeContent string

type errMsg error

type model struct {
	config *config.Context
	err    errMsg

	width              int       // the terminal max width
	height             int       // the terminal max height
	mode               mode.Mode // current mode
	showWelcomeContent bool
	showPreview        bool

	writeArea   textarea.Model
	previewView *preview.View
	commandLine *commandline.Bar
	statusLine  *statueline.Bar

	md *doc.Markdown // current edit markdown file
}

// New ui
func New(config *config.Context) *model {
	m := model{}
	m.config = config
	m.showWelcomeContent = config.ShouldShowWelcomeContent()
	m.mode = mode.Normal
	m.md = doc.DefaultMd()
	m.previewView = preview.New(config)
	m.statusLine = statueline.New(config)
	m.commandLine = commandline.New(config)
	return &m
}

func (m *model) Init() tea.Cmd {
	m.writeArea = textarea.New()
	m.writeArea.KeyMap = m.config.Keymap.InsertModeKeyMap
	m.writeArea.CharLimit = -1

	if m.showWelcomeContent {
		m.md = doc.DefaultMd()
		m.previewView.SetContent(welcomeContent)
		return nil
	} else {
		m.md, m.err = doc.FilePathToMd(m.config.InitFilePath)
		m.writeArea.SetValue(m.md.Body)
		m.previewView.SetContent(m.md.Body)
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
			m.md.FileName = msg.Value()
			m.savefile()
			m.refreshStatusLine()
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.setsize(msg)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.config.Keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.config.Keymap.SaveFile):
			if m.showWelcomeContent {
				return m, nil
			}

			if m.md.NoName() {
				return m, m.prompt(&command.SaveFile{}) // todo 使用 modeCommand 模式 获取输入内容
			}
			m.savefile()
			return m, nil
		case key.Matches(msg, m.config.Keymap.PreviewView):
			m.showPreview = !m.showPreview
			return m, nil
		case key.Matches(msg, m.config.Keymap.ToCommandMode):
			//if m.mode == modeNormal && !m.showWelcomeContent {
			//	m.toCommandMode()
			//	return m, m.Bar.Focus()
			//}
			// allow in welcome View use command mode
			m.toCommandMode()
			return m, m.commandLine.Focus()
		case key.Matches(msg, m.config.Keymap.ToNormalMode):
			m.toNormalMode()
		case key.Matches(msg, m.config.Keymap.ToInsertMode):
			if m.mode == mode.Normal {
				m.toInsertMode()
				return m, nil
			}
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	switch m.mode {
	case mode.Insert:
		m.writeArea, cmd = m.writeArea.Update(msg)
		cmds = append(cmds, cmd)

		m.previewView.SetContent(m.writeArea.Value())
	case mode.Command:
		cmd = m.commandLine.Update(msg)
		cmds = append(cmds, cmd)
	}

	if !m.showWelcomeContent {
		m.refreshStatusLine()
	}

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	m.resize()
	if m.showWelcomeContent {
		return m.welcomeContentView()
	}

	buffer := utils.NewStrBuffer()

	if m.showPreview { // todo 当前只是简单的处理, 默认是左右视图
		if m.width >= 160 {
			val := lipgloss.JoinHorizontal(lipgloss.Top,
				lipgloss.NewStyle().AlignHorizontal(lipgloss.Right).Render(utils.TimeIt(m.writeArea.View)),
				lipgloss.NewStyle().AlignHorizontal(lipgloss.Left).Render(m.previewView.View()),
			)
			buffer.Write(val)
		} else {
			buffer.Write(m.previewView.View())
		}
	} else {
		buffer.Write(m.writeArea.View())
	}

	return buffer.
		NewLine().
		Write(m.statusLine.View()).
		NewLine().
		Write(m.commandLine.View()).
		String()
}

func (m *model) resize() {
	set := func(width, height int) {
		m.previewView.Set(width, height)
		m.writeArea.SetWidth(width)
		m.writeArea.SetHeight(height)
	}

	midWidth := m.width / 2
	if m.showPreview && m.width >= 160 {
		set(midWidth, m.height-2)
	} else {
		set(m.width, m.height-2)
	}
}

func (m *model) setsize(msg tea.WindowSizeMsg) {
	m.width = msg.Width
	m.height = msg.Height
}

func (m *model) toCommandMode() {
	m.mode = mode.Command
	m.commandLine.Show()
	m.refreshStatusLine()
}

func (m *model) toInsertMode() {
	m.mode = mode.Insert
	if m.showWelcomeContent {
		m.showWelcomeContent = false
		m.previewView.SetContent("")
	}

	m.writeArea.Focus()
	m.refreshStatusLine()
}

func (m *model) toNormalMode() {
	m.mode = mode.Normal
	m.writeArea.Blur()
	m.commandLine.Hide()
}

func (m *model) prompt(cmd command.Command) tea.Cmd {
	m.mode = mode.Command
	m.writeArea.Blur()
	m.refreshStatusLine()
	return m.commandLine.Prompt(cmd)
}

func (m *model) hideCommandLine() {
	m.commandLine.Hide()
}

func (m *model) refreshStatusLine() {
	area := m.writeArea
	m.statusLine.Refresh(m.md, m.width, m.height, m.mode, area.Line(), area.LineInfo().ColumnOffset, area.LineCount())
}

func (m *model) savefile() {
	m.md.Body = m.writeArea.Value()
	m.err = m.md.Save()
}

func (m *model) welcomeContentView() string {
	m.writeArea.View()
	return utils.NewStrBuffer().
		Write(m.previewView.View()).
		NewLine().
		Write(m.commandLine.View()).
		String()
}
