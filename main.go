package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/md/config"
	"github.com/fzdwx/md/ui"
	"os"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	cfg := config.Parse()

	p := tea.NewProgram(ui.New(cfg), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
