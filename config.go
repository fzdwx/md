package main

import (
	"os"
)

type mdConfig struct {
	filePath string
	style    string
}

func parseConfig() *mdConfig {
	args := os.Args
	var cfg mdConfig
	if len(args) > 1 {
		cfg.filePath = args[1]
	}

	cfg.style = "dark"

	return &cfg
}

// showWelcomeContent if user don't choose markdown file then show welcome content
func (c mdConfig) showWelcomeContent() bool {
	return c.filePath == ""
}
