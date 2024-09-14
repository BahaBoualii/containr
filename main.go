package main

import (
	"fmt"
	"os"

	"github.com/BahaBoualii/containr/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := tui.New()
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
