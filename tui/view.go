package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	headerStyle   = lipgloss.NewStyle().
			Bold(true).                        // Make text bold
			Foreground(lipgloss.Color("212")). // Dark gray text color
			Background(lipgloss.Color("66")).  // Light blue background
			Padding(1, 2).                     // Add padding around text
			Margin(1, 0).                      // Add margin around the header
			Border(lipgloss.NormalBorder())    // Add a border
)

func (m model) View() string {
	header := headerStyle.Render("Containr")
	s := header + "\n\n"
	for i, item := range m.choices {
		if m.cursor == i {
			s += selectedStyle.Render(fmt.Sprintf("> %s", item)) + "\n"
		} else {
			s += normalStyle.Render(fmt.Sprintf("  %s", item)) + "\n"
		}
	}
	s += "\nPress Enter to select, Up/Down to navigate."
	return s
}
