package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Quit              key.Binding       // quit program
	ToNormalMode      key.Binding       // to normal mode
	ToCommandMode     key.Binding       // to normal mode
	ToInsertMode      key.Binding       // to insert mode , edit file
	insertModeKeyMap  textarea.KeyMap   // insert mode key map(write area key map)
	commandLineKeyMap CommandLineKeyMap // command line mode key map

	SaveFile key.Binding // save file to disk
}

type CommandLineKeyMap struct {
	Cr key.Binding // same vim <CR>,  common is Enter
}

var DefaultKeyMap = KeyMap{
	Quit:             key.NewBinding(key.WithKeys(tea.KeyCtrlC.String())),
	ToNormalMode:     key.NewBinding(key.WithKeys(tea.KeyEsc.String())),
	ToInsertMode:     key.NewBinding(key.WithKeys("i")),
	ToCommandMode:    key.NewBinding(key.WithKeys(":")),
	insertModeKeyMap: textarea.DefaultKeyMap,
	commandLineKeyMap: CommandLineKeyMap{
		Cr: key.NewBinding(key.WithKeys(tea.KeyEnter.String())),
	},
	SaveFile: key.NewBinding(key.WithKeys(tea.KeyCtrlS.String())),
}
