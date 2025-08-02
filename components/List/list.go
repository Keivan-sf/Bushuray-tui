package list

import (
	"fmt"
	"log"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type ListItem struct {
	Name       string
	Protocol   string
	TestResult int
}

type Model struct {
	Id      string
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
	log.Println("got update on list")
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if l.cursor > 0 {
				l.cursor--
			}
			adjustOffsetForCursor(&l)
		case "down", "j":
			if l.cursor < len(l.Items)-1 {
				l.cursor++
			}
			adjustOffsetForCursor(&l)
		case "enter":
			l.Primary = l.cursor
		}
	case tea.MouseMsg:
		switch msg.Button {
		case tea.MouseButtonLeft:
			for i := range l.Items {
				if zone.Get(l.Id + strconv.Itoa(i)).InBounds(msg) {
					l.cursor = i
				}
			}
		case tea.MouseButtonWheelDown:
			if !zone.Get(l.Id).InBounds(msg) {
				break
			}
			if l.offset < len(l.Items)-l.Height {
				l.offset++
			}
			adjustCursorForOffset(&l)
		case tea.MouseButtonWheelUp:
			if !zone.Get(l.Id).InBounds(msg) {
				break
			}
			if l.offset > 0 {
				l.offset--
			}
			adjustCursorForOffset(&l)
		}
	}

	return l, nil
}

func (l Model) View() string {
	var s []string
	start := l.offset
	end := len(l.Items)
	w := l.Width - protocol_w - test_result_w
	for i := start; i < end; i++ {
		row := ""
		item := l.Items[i]
		item_str := fmt.Sprintf(" %s", item.Name)
		if i == l.Primary {
			protocol := protocol_primary_style.Render(item.Protocol)
			row += protocol
			row += primary_style.Width(w).MaxWidth(w).MaxHeight(1).Render(item_str) + styleTestPrimary(item.TestResult)
		} else if i == l.cursor {
			protocol := protocol_under_cursor_style.Render(item.Protocol)
			row += protocol
			row += under_cursor_style.Width(w).MaxWidth(w).MaxHeight(1).Render(item_str) + styleTestUnderCursor(item.TestResult)
		} else {
			protocol := protocol_style.Render(item.Protocol)
			row += protocol
			row += item_style.Width(w).MaxWidth(w).MaxHeight(1).Render(item_str) + styleTestNormal(item.TestResult)
		}
		s = append(s, zone.Mark(l.Id+strconv.Itoa(i), row))
	}
	list_view := zone.Mark(l.Id, lipgloss.NewStyle().Height(l.Height).MaxHeight(l.Height).Render(lipgloss.JoinVertical(lipgloss.Top, s...)))
	return list_view
}
