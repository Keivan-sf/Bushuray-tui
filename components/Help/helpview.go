package helpview

import (
	cmds "bushuray-tui/commands"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyHelp struct {
	Key  string
	Help string
}

type Model struct {
	Width  int
	Height int
	keys   []KeyHelp
}

func InitialModel() Model {
	return Model{
		keys: []KeyHelp{
			{Key: "enter", Help: "connect"},
			{Key: "p", Help: "paste config(s)"},
			{Key: "v", Help: "toggle tun mode"},
			{Key: "a", Help: "add a new group"},
			{Key: "t", Help: "test config"},
			{Key: "T", Help: "test group configs"},
			{Key: "del/d", Help: "delete config"},
			{Key: "D", Help: "delete group"},
			{Key: "k/↑", Help: "move up"},
			{Key: "j/↓", Help: "move down"},
			{Key: "tab", Help: "next tab"},
			{Key: "ctrl+pgdn", Help: "next tab"},
			{Key: "shift+tab", Help: "previous tab"},
			{Key: "ctrl+pgup", Help: "previous tab"},
			{Key: "ctrl+c", Help: "exit"},
			{Key: "?", Help: "help menu"},
		},
	}
}

func (m Model) View() string {
	return m.GenHelpKey(m.keys)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q", "?":
			return m, cmds.ExitHelpView
		}
	}
	return m, nil
}

func (m Model) SetWH(width int, height int) Model {
	m.Width = width
	m.Height = height
	return m

}

var secondary_style = lipgloss.NewStyle().Foreground(lipgloss.Color("#4c4f69"))
var primary_style = lipgloss.NewStyle()

func (m Model) GenHelpKey(keys []KeyHelp) string {
	key_max_width := 0
	for _, kh := range keys {
		key_width := lipgloss.Width(kh.Key)
		if key_width > key_max_width {
			key_max_width = key_width
		}
	}

	var rows []string
	for _, kh := range keys {
		key_str := primary_style.Width(key_max_width + 2).Render(kh.Key)
		help_str := secondary_style.Render(kh.Help)
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Left, key_str, help_str))
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)
	container := lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, content)
	return container
}
