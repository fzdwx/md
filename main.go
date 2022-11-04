package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

func main() {
	open, err := os.OpenFile("qwe.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	log.SetOutput(open)
	if err != nil {
		panic(err)
	}

	config := parseConfig()
	p := tea.NewProgram(initialModel(config))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
