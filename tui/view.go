package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.loaded {
		containersView := m.lists[containerList].View()
		imagesView := m.lists[imageList].View()
		volumesView := m.lists[volumeList].View()
		switch m.focused {
		case imageList:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(containersView),
				focusedStyle.Render(imagesView),
				columnStyle.Render(volumesView))
		case volumeList:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(containersView),
				columnStyle.Render(imagesView),
				focusedStyle.Render(volumesView))
		default: // containerList
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
