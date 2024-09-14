package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.loaded {
		containersView := m.lists[containers].View()
		imagesView := m.lists[images].View()
		volumesView := m.lists[volumes].View()
		switch m.focused {
		case images:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(containersView),
				focusedStyle.Render(imagesView),
				columnStyle.Render(volumesView))
		case volumes:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(containersView),
				columnStyle.Render(imagesView),
				focusedStyle.Render(volumesView))
		default:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				focusedStyle.Render(containersView),
				columnStyle.Render(imagesView),
				columnStyle.Render(volumesView))
		}
	} else {
		return "Loading..."
	}
}
