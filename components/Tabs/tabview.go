package tabs

import (
	"strings"

	list "github.com/Keivan-sf/Bushuray-tui/components/List"
	"github.com/Keivan-sf/Bushuray-tui/global"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TabView struct {
	Title   string
	Content list.Model
}

var border_color = lipgloss.Color("#7287fd")
var title_color = lipgloss.Color("#4c4f69")

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
		Foreground(title_color).
		Background(global.GetBgColor()).
		BorderBackground(global.GetBgColor()).
		BorderForeground(border_color).
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
		Background(global.GetBgColor()).
		BorderBackground(global.GetBgColor()).
		Border(tabBorder, true).
		BorderForeground(border_color).
		Padding(0, 1).Render(title)
	return tab
}

func renderTabLine(width int) string {
	extra_line_str := strings.Repeat(" ", width)
	extra_line := lipgloss.NewStyle().
		Border(lipgloss.Border{Bottom: "─", BottomLeft: "─"}).
		Background(global.GetBgColor()).
		BorderBackground(global.GetBgColor()).
		BorderForeground(border_color).
		Render(extra_line_str)
	return extra_line
}
