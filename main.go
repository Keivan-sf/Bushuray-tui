package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ListItem struct {
	Name string
}

type List struct {
	Items  []ListItem
	cursor int
	width  int
	height int
}

var selected_style = lipgloss.NewStyle().Background(lipgloss.Color("#ea76cb"))
var under_cursor_style = lipgloss.NewStyle().Background(lipgloss.Color("#8839ef"))

func (l List) Update(msg tea.Msg) (List, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if l.cursor > 0 {
				l.cursor--
			}
			return l, nil
		case "down", "j":
			if l.cursor < len(l.Items)-1 {
				l.cursor++
			}
			return l, nil
		}

	}
	return l, nil
}

func (l List) View() string {
	s := ""
	for i, item := range l.Items {
		item_str := fmt.Sprintf("%d - %s", i, item.Name)
		if i == l.cursor {
			s += under_cursor_style.Width(l.width).Render(item_str) + "\n"
		} else {
			s += item_str + "\n"
		}
	}
	return s
}

type Model struct {
	ConfigList List
	width      int
	height     int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ConfigList.width = msg.Width
		m.ConfigList.height = msg.Height
	}

	var cmd tea.Cmd
	m.ConfigList, cmd = m.ConfigList.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.ConfigList.View()
}

func initModel() Model {
	return Model{
		ConfigList: List{Items: []ListItem{{Name: "Hello"}, {Name: "There"}, {Name: "Friend"}, {Name: "How"}, {Name: "Are"}, {Name: "You"}}},
	}
}

func main() {
	p := tea.NewProgram(initModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}
}
