package main

import (
	"flag"
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var cpuprofile = "cpu.f"
var memprofile = "mem.f"

func main() {
	flag.Parse()
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	open, err := os.OpenFile("qwe.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	log.SetOutput(open)
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

var welcomeContent = "# hello world\n zzzzzzzzzzz"

type errMsg error

type mode int

const (
	welcome mode = iota + 1
	normal
	insert
)

type model struct {
	showWelcomeContent bool

	mode            mode
	writeArea       textarea.Model
	welcomeViewPort viewport.Model
	err             errMsg
}

// initialModel
func initialModel() *model {
	m := model{mode: normal}
	m.showWelcomeContent = true
	return &m
}

func toMd(in string, stylePath string) (string, error) {
	return glamour.Render(in, stylePath)
}

func mustToMd(in string, stylePath string) string {
	md, _ := toMd(in, stylePath)
	return md
}

func (m *model) Init() tea.Cmd {
	if m.showWelcomeContent {
		m.mode = welcome
		m.welcomeViewPort = viewport.New(10, 10)
		m.welcomeViewPort.SetContent(mustToMd(welcomeContent, "dark"))
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
		return m.welcomeViewPort.View()
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
