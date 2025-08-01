package list

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ListItem struct {
	Name       string
	Protocol   string
	TestResult int
}

type Model struct {
	Items   []ListItem
	cursor  int
	Width   int
	Height  int
	offset  int
	Primary int
}

// var primary_style = lipgloss.NewStyle().Background(lipgloss.Color("#4c4f69"))
var primary_style = lipgloss.NewStyle().Background(lipgloss.Color("#4c4f69"))

// var protocol_primary_style = lipgloss.NewStyle().Background(lipgloss.Color("#5c5f77")).Foreground(lipgloss.Color("#9ca0b0")).Width(8).Align(lipgloss.Center)
var protocol_primary_style = lipgloss.NewStyle().Background(lipgloss.Color("#40a02b")).Foreground(lipgloss.Color("#FFF")).Width(8).Align(lipgloss.Center)
var under_cursor_style = lipgloss.NewStyle().Background(lipgloss.Color("#1e2030"))
var protocol_under_cursor_style = lipgloss.NewStyle().Background(lipgloss.Color("#24273a")).Foreground(lipgloss.Color("#9ca0b0")).Width(8).Align(lipgloss.Center)
var protocol_style = lipgloss.NewStyle().Foreground(lipgloss.Color("#4c4f69")).Width(8).Align(lipgloss.Center)

func (l Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if l.cursor > 0 {
				l.cursor--
			}
		case "down", "j":
			if l.cursor < len(l.Items)-1 {
				l.cursor++
			}
		case "enter":
			l.Primary = l.cursor
		}

	}
	if l.cursor < l.offset {
		l.offset = l.cursor
	} else if l.cursor >= l.offset+l.Height {
		l.offset = l.cursor - l.Height + 1
	}
	return l, nil
}

func (l Model) View() string {
	s := ""
	start := l.offset
	end := len(l.Items)
	for i := start; i < end; i++ {
		item := l.Items[i]
		item_str := fmt.Sprintf(" %s", item.Name)
		if i == l.Primary {
			protocol := protocol_primary_style.Render(item.Protocol)
			s += protocol
			s += primary_style.Width(l.Width-7).Render(item_str) + "\n"
		} else if i == l.cursor {
			protocol := protocol_under_cursor_style.Render(item.Protocol)
			s += protocol
			s += under_cursor_style.Width(l.Width-7).Render(item_str) + "\n"
		} else {
			protocol := protocol_style.Render(item.Protocol)
			s += protocol
			s += item_str + "\n"
		}
	}
	s = lipgloss.NewStyle().Height(l.Height).MaxHeight(l.Height).Render(s)
	return s
}
