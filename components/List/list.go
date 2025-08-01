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

var protocol_w = 8
var test_result_w = 15

var primary_style = lipgloss.NewStyle().Background(lipgloss.Color("#4c4f69"))
var protocol_primary_style = lipgloss.NewStyle().Background(lipgloss.Color("#40a02b")).Foreground(lipgloss.Color("#FFF")).Width(protocol_w).Align(lipgloss.Center)

var under_cursor_style = lipgloss.NewStyle().Background(lipgloss.Color("#1e2030"))
var protocol_under_cursor_style = lipgloss.NewStyle().Background(lipgloss.Color("#24273a")).Foreground(lipgloss.Color("#9ca0b0")).Width(protocol_w).Align(lipgloss.Center)

var protocol_style = lipgloss.NewStyle().Foreground(lipgloss.Color("#4c4f69")).Width(protocol_w).Align(lipgloss.Center)
var item_style = lipgloss.NewStyle()

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
	w := l.Width - protocol_w - test_result_w
	for i := start; i < end; i++ {
		item := l.Items[i]
		item_str := fmt.Sprintf(" %s", item.Name)
		if i == l.Primary {
			protocol := protocol_primary_style.Render(item.Protocol)
			s += protocol
			s += primary_style.Width(w).MaxWidth(w).Render(item_str) + styleTestPrimary(item.TestResult) + "\n"
		} else if i == l.cursor {
			protocol := protocol_under_cursor_style.Render(item.Protocol)
			s += protocol
			s += under_cursor_style.Width(w).MaxWidth(w).Render(item_str) + styleTestUnderCursor(item.TestResult) + "\n"
		} else {
			protocol := protocol_style.Render(item.Protocol)
			s += protocol
			s += item_style.Width(w).MaxWidth(w).Render(item_str) + styleTestNormal(item.TestResult) + "\n"
		}
	}
	s = lipgloss.NewStyle().Height(l.Height).MaxHeight(l.Height).Render(s)
	return s
}
