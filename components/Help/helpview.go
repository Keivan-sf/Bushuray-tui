package helpview

import (
	cmds "bushuray-tui/commands"
	"bushuray-tui/global"

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
			{Key: "p/ctrl+v", Help: "paste profile(s)"},
			{Key: "y", Help: "copy profile"},
			{Key: "v", Help: "toggle tun mode"},
			{Key: "a", Help: "add a new group"},
			{Key: "t", Help: "test profile"},
			{Key: "T", Help: "test group profiles"},
			{Key: "U", Help: "update subscription"},
			{Key: "S", Help: "sort by test result"},
			{Key: "del/d", Help: "delete profile"},
			{Key: "D", Help: "delete group"},
			{Key: "k/↑", Help: "move up"},
			{Key: "j/↓", Help: "move down"},
			{Key: "J", Help: "jump to active profile"},
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
	help_max_width := 0
	for _, kh := range keys {
		key_width := lipgloss.Width(kh.Key)
		help_width := lipgloss.Width(kh.Help)
		if key_width > key_max_width {
			key_max_width = key_width
		}
		if help_width > help_max_width {
			help_max_width = help_width
		}
	}

	var rows []string
	bg_style := lipgloss.NewStyle().Background(global.GetBgColor()).Width(m.Width).Height(1).Align(lipgloss.Center)
	help_row_style := bg_style.Width(help_max_width + key_max_width + 2).Align(lipgloss.Left)
	for _, kh := range keys {
		key_str := primary_style.Background(global.GetBgColor()).Width(key_max_width + 2).Render(kh.Key)
		help_str := secondary_style.Background(global.GetBgColor()).Render(kh.Help)
		row_content := help_row_style.Render(lipgloss.JoinHorizontal(lipgloss.Left, key_str, help_str))
		row := bg_style.Render(row_content)
		rows = append(rows, row)
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)
	container := lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, content)
	container_with_bg := lipgloss.NewStyle().Background(global.GetBgColor()).Render(container)
	return container_with_bg
}
