package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/md/config"
	"github.com/fzdwx/md/ui"
	"os"
)

func main() {
	cfg := config.Parse()

	p := tea.NewProgram(ui.New(cfg), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
