package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	columnStyle = lipgloss.NewStyle().
			Padding(1, 2)
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
)

type category int

const (
	containers category = iota
	images
	volumes
)

type DockerOption struct {
	category    category
	title       string
	description string
}

func (t DockerOption) FilterValue() string {
	return t.title
}

func (t DockerOption) Title() string {
	return t.title
}

func (t DockerOption) Description() string {
	return t.description
}

type Model struct {
	lists   []list.Model
	focused category
	loaded  bool
}

func New() *Model {
	return &Model{}
}

func (m *Model) Next() {
	if m.focused == volumes {
		m.focused = containers
	} else {
		m.focused++
	}
}

func (m *Model) Previous() {
	if m.focused == containers {
		m.focused = volumes
	} else {
		m.focused--
	}
}

func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/4, height-4)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	m.lists[containers].Title = "Containers"
	m.lists[containers].SetItems([]list.Item{
		DockerOption{category: containers, title: "List Containers", description: "Get all containers info"},
		DockerOption{category: containers, title: "Run Container", description: "Run a container from a specific image"},
	})

	m.lists[images].Title = "Images"
	m.lists[images].SetItems([]list.Item{
		DockerOption{category: images, title: "List Images", description: "Get all images info"},
		DockerOption{category: images, title: "Remove Image", description: "Remove a specific image"},
	})

	m.lists[volumes].Title = "volumes"
	m.lists[volumes].SetItems([]list.Item{
		DockerOption{category: volumes, title: "List Volumes", description: "Get all volumes info"},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			columnStyle.Width(msg.Width / 4)
			focusedStyle.Width(msg.Width / 4)
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left", "h":
			m.Previous()
		case "right", "l":
			m.Next()
		}
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

// func (m Model) handleChoice() model {
// 	switch m.choices[m.cursor] {
// 	case "List Containers":
// 		m = m.listContainers()
// 	case "Start Container":
// 		// Implement logic to start container
// 		m = m.startContainer()
// 	case "Stop Container":
// 		// Implement logic to stop container
// 		m = m.stopContainer()
// 	case "Pull Image":
// 		// Implement logic to pull image
// 	}

// 	return m
// }

// func (m model) listContainers() model {
// 	containerDetails, err := containers.ListAllContainers()
// 	if err != nil {
// 		m.choices = []string{err.Error()}
// 		return m
// 	}
// 	m.choices = containerDetails
// 	return m
// }

// func (m model) startContainer() model {
// 	// Assume the first container in the list is selected for simplicity.
// 	// In a real-world scenario, you would select a container ID from the user input.
// 	if len(m.choices) > 0 {
// 		containerID := m.choices[m.cursor][:10]
// 		err := containers.StartContainer(containerID)
// 		if err != nil {
// 			m.choices = []string{err.Error()}
// 		} else {
// 			m.choices = []string{"Container started successfully"}
// 		}
// 	}

// 	return m
// }

// func (m model) stopContainer() model {
// 	// Same logic as startContainer for stopping a container
// 	if len(m.choices) > 0 {
// 		containerID := m.choices[m.cursor][:10]
// 		err := containers.StopContainer(containerID)
// 		if err != nil {
// 			m.choices = []string{err.Error()}
// 		} else {
// 			m.choices = []string{"Container stopped successfully"}
// 		}
// 	}

// 	return m
// }

// func (m model) removeContainer() model {
// 	// Same logic as startContainer for removing a container
// 	if len(m.choices) > 0 {
// 		containerID := m.choices[m.cursor][:10]
// 		err := containers.RemoveContainer(containerID)
// 		if err != nil {
// 			m.choices = []string{err.Error()}
// 		} else {
// 			m.choices = []string{"Container removed successfully"}
// 		}
// 	}

// 	return m
// }
