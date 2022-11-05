package main

import (
	"os"
)

type mdConfig struct {
	filePath string
	mdStyle  string

	keymap KeyMap
}

func parseConfig() *mdConfig {
	args := os.Args
	var cfg mdConfig
	if len(args) > 1 {
		cfg.filePath = args[1]
	}

	cfg.mdStyle = "dark"

	cfg.keymap = DefaultKeyMap
	return &cfg
}

// showWelcomeContent if user don't choose markdown file then show welcome content
func (c mdConfig) showWelcomeContent() bool {
	return c.filePath == ""
}
