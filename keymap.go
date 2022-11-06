package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Quit             key.Binding     // quit program
	ToNormalMode     key.Binding     // to normal mode
	ToCommandMode    key.Binding     // to normal mode
	ToInsertMode     key.Binding     // to insert mode , edit file
	insertModeKeyMap textarea.KeyMap // insert mode key map(write area key map)

	SaveFile key.Binding // save file to disk
}

var DefaultKeyMap = KeyMap{
	Quit:             key.NewBinding(key.WithKeys(tea.KeyCtrlC.String())),
	ToNormalMode:     key.NewBinding(key.WithKeys(tea.KeyEsc.String())),
	ToInsertMode:     key.NewBinding(key.WithKeys("i")),
	ToCommandMode:    key.NewBinding(key.WithKeys(":")),
	insertModeKeyMap: textarea.DefaultKeyMap,

	SaveFile: key.NewBinding(key.WithKeys(tea.KeyCtrlS.String())),
}
