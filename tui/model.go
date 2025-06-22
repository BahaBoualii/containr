package tui

import (
	"github.com/BahaBoualii/containr/pkg/containers"
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

type listCategory int

const (
	containerList listCategory = iota
	imageList
	volumeList
)

type DockerOption struct {
	category    listCategory
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
	focused listCategory
	loaded  bool
}

func New() *Model {
	return &Model{}
}

func (m *Model) Next() {
	if m.focused == volumeList {
		m.focused = containerList
	} else {
		m.focused++
	}
}

func (m *Model) Previous() {
	if m.focused == containerList {
		m.focused = volumeList
	} else {
		m.focused--
	}
}

func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/4, height-4)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	m.lists[containerList].Title = "Containers"
	m.lists[containerList].SetItems([]list.Item{
		DockerOption{category: containerList, title: "List Containers", description: "Get all containers info"},
		DockerOption{category: containerList, title: "Run Container", description: "Run a container from a specific image"},
	})

	m.lists[imageList].Title = "Images"
	m.lists[imageList].SetItems([]list.Item{
		DockerOption{category: imageList, title: "List Images", description: "Get all images info"},
		DockerOption{category: imageList, title: "Remove Image", description: "Remove a specific image"},
	})

	m.lists[volumeList].Title = "volumes"
	m.lists[volumeList].SetItems([]list.Item{
		DockerOption{category: volumeList, title: "List Volumes", description: "Get all volumes info"},
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
		case "enter":
			// Only handle choice if the focused list is the initial options list
			if m.lists[m.focused].SelectedItem() != nil {
				option, ok := m.lists[m.focused].SelectedItem().(DockerOption)
				if ok && option.description != "" { // Check if it's an actionable item with a description
					m.handleChoice()
				}
			}
		case "s":
			// Only stop container if the focused list is containers and an item is selected
			if m.focused == containerList && m.lists[containerList].SelectedItem() != nil {
				// Check if the items are actual containers (not the initial options)
				// This can be done by checking if description is empty or a specific value
				// For now, let's assume if description is empty, it's a listed container
				selectedItem, ok := m.lists[containerList].SelectedItem().(DockerOption)
				if ok && selectedItem.description == "" {
					m.stopContainer()
				}
			}
		}
	}
	var cmd tea.Cmd
	if m.loaded {
		m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	}
	return m, cmd
}

func (m *Model) handleChoice() {
	selectedItem, ok := m.lists[m.focused].SelectedItem().(DockerOption)
	if !ok {
		return
	}

	switch selectedItem.title {
	case "List Containers":
		m.listContainers()
		// case "Start Container":
		// 	// Implement logic to start container
		// 	m = m.startContainer()
		// case "Stop Container":
		// 	// Implement logic to stop container
		// 	m = m.stopContainer()
		// case "Pull Image":
		// 	// Implement logic to pull image
	}
}

func (m *Model) listContainers() {
	containerDetails, err := containers.ListAllContainers()
	if err != nil {
		// Handle error, maybe show a message to the user
		// For now, we'll just keep the existing items
		return
	}

	var items []list.Item
	for _, detail := range containerDetails {
		items = append(items, DockerOption{category: containerList, title: detail, description: ""}) // Assuming description is not available or needed for now
	}
	m.lists[containerList].SetItems(items)
}

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

func (m *Model) stopContainer() {
	selectedItem, ok := m.lists[m.focused].SelectedItem().(DockerOption)
	if !ok {
		// Should not happen if an item is selected
		return
	}

	// Assuming the title is in the format "ID: Image"
	containerID := selectedItem.title[:10] // Extract first 10 chars as ID

	err := containers.StopContainer(containerID)
	if err != nil {
		// Handle error, maybe show a message to the user
		// For now, we'll just list containers again to refresh
		m.listContainers()
		return
	}

	// Refresh the list after stopping
	m.listContainers()
	// Optionally, set a success message to be displayed in the View
	// m.message = "Container stopped successfully"
}

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
