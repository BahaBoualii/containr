package tui

import (
	"github.com/BahaBoualii/containr/pkg/containers"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string // items on the to-do list
	cursor   int      // which to-do list item our cursor is pointing at
	selected map[int]struct{}
}

func InitialModel() model {
	return model{
		choices:  []string{"List Containers", "Start Container", "Stop Container", "Pull Image"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				m = m.handleChoice()
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) handleChoice() model {
	switch m.choices[m.cursor] {
	case "List Containers":
		m = m.listContainers()
	case "Start Container":
		// Implement logic to start container
		m = m.startContainer()
	case "Stop Container":
		// Implement logic to stop container
		m = m.stopContainer()
	case "Pull Image":
		// Implement logic to pull image
	}

	return m
}

func (m model) listContainers() model {
	containerDetails, err := containers.ListAllContainers()
	if err != nil {
		m.choices = []string{err.Error()}
		return m
	}
	m.choices = containerDetails
	return m
}

func (m model) startContainer() model {
	// Assume the first container in the list is selected for simplicity.
	// In a real-world scenario, you would select a container ID from the user input.
	if len(m.choices) > 0 {
		containerID := m.choices[m.cursor][:10]
		err := containers.StartContainer(containerID)
		if err != nil {
			m.choices = []string{err.Error()}
		} else {
			m.choices = []string{"Container started successfully"}
		}
	}

	return m
}

func (m model) stopContainer() model {
	// Same logic as startContainer for stopping a container
	if len(m.choices) > 0 {
		containerID := m.choices[m.cursor][:10]
		err := containers.StopContainer(containerID)
		if err != nil {
			m.choices = []string{err.Error()}
		} else {
			m.choices = []string{"Container stopped successfully"}
		}
	}

	return m
}

func (m model) removeContainer() model {
	// Same logic as startContainer for removing a container
	if len(m.choices) > 0 {
		containerID := m.choices[m.cursor][:10]
		err := containers.RemoveContainer(containerID)
		if err != nil {
			m.choices = []string{err.Error()}
		} else {
			m.choices = []string{"Container removed successfully"}
		}
	}

	return m
}
