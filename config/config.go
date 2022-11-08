package config

import (
	"github.com/fzdwx/md/config/keymap"
	"github.com/fzdwx/md/config/style"
	"os"
)

type Context struct {
	InitFilePath string
	MdStyle      string

	Keymap keymap.KeyMap
	Style  style.Style
}

func Parse() *Context {
	args := os.Args
	var cfg Context
	if len(args) > 1 {
		cfg.InitFilePath = args[1]
	}

	cfg.MdStyle = "dark"

	cfg.Keymap = keymap.DefaultKeyMap
	cfg.Style = style.DefaultStyle
	return &cfg
}

// ShouldShowWelcomeContent if user don't choose markdown file then Show welcome content
func (c Context) ShouldShowWelcomeContent() bool {
	return c.InitFilePath == ""
}
