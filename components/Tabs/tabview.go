package tabs

import (
	"bushuray-tui/components/List"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TabView struct {
	Title   string
	Content list.Model
}

func (t TabView) View() string {
	return t.Content.View()
}

func (t TabView) Update(msg tea.Msg) (TabView, tea.Cmd) {
	var cmd tea.Cmd
	t.Content, cmd = t.Content.Update(msg)
	return t, cmd
}

func (t TabView) SetWH(width int, height int) TabView {
	t.Content.Height = height
	t.Content.Width = width
	return t
}

func renderTitle(title string) string {
	tabBorder := lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab := lipgloss.NewStyle().
		Border(tabBorder, true).
		Padding(0, 1).Render(title)
	return tab
}

func renderActiveTitle(title string) string {
	tabBorder := lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tab := lipgloss.NewStyle().
		Border(tabBorder, true).
		Padding(0, 1).Render(title)
	return tab
}
