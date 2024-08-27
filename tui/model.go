package tui

import (
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
	// Handle user inputs and update state
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			m = m.handleChoice()
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m model) handleChoice() model {
	switch m.choices[m.cursor] {
	case "List Containers":
		m.listContainers()
	case "Start Container":
		// Implement logic to start container
	case "Stop Container":
		// Implement logic to stop container
	case "Pull Image":
		// Implement logic to pull image
	}

	return m
}

func (m model) listContainers() {
	// cli, err := client.NewClientWithOpts(client.FromEnv)
	// if err != nil {
	// 	m.choices = []string{"Error: unable to connect to Docker"}
	// 	return
	// }

	// containers, err := cli.ContainerList(
	// 	context.Background(),
	// 	container.ListOptions{All: true})

	// if err != nil {
	// 	m.choices = []string{"Error: unable to list containers"}
	// 	return
	// }

	m.choices = []string{"hello1", "hello2"}
	// for _, container := range containers {
	// 	m.choices = append(m.choices, fmt.Sprintf("%s: %s", container.ID[:10], container.Image))
	// }
}
