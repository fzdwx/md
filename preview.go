package main

import "github.com/charmbracelet/glamour"

func toMd(in string, stylePath string) (string, error) {
	return glamour.Render(in, stylePath)
}

func mustToMd(in string, stylePath string) string {
	md, _ := toMd(in, stylePath)
	return md
}
