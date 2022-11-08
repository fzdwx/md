package style

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"testing"
)

func Test_Color(t *testing.T) {
	fmt.Println(lipgloss.NewStyle().Background(lipgloss.Color(A1)).Render("hello world"))
	fmt.Println(lipgloss.NewStyle().Background(lipgloss.Color(A2)).Render("hello world"))
	fmt.Println(lipgloss.NewStyle().Background(lipgloss.Color(B1)).Render("hello world"))
	fmt.Println(lipgloss.NewStyle().Background(lipgloss.Color(B2)).Render("hello world"))
	fmt.Println(lipgloss.NewStyle().Background(lipgloss.Color(B3)).Render("hello world"))
}
