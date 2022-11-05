package main

import (
	"fmt"
	"github.com/charmbracelet/glamour"
	"io"
	"os"
)

func main() {
	r, err := os.Open("example.md")
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	out, err := glamour.Render(string(bytes), "dark")
	if err != nil {
		panic(err)
	}

	fmt.Print(out)
}
