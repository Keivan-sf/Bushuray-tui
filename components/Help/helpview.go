package helpview

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Width  int
	Height int
	keys   []KeyHelp
}

type KeyHelp struct {
	key  string
	help string
}

func InitialModel() Model {
	return Model{
		keys: []KeyHelp{
			{key: "enter", help: "connect"},
			{key: "p", help: "paste config(s)"},
			{key: "v", help: "toggle tun mode"},
			{key: "a", help: "add a new group"},
			{key: "t", help: "test config"},
			{key: "T", help: "test group configs"},
			{key: "del/d", help: "delete config"},
			{key: "D", help: "delete group"},
			{key: "k/↑", help: "move up"},
			{key: "j/↓", help: "move down"},
			{key: "tab/ctrl+pgdn", help: "next tab"},
			{key: "shift+tab/ctrl+pgup", help: "previous tab"},
			{key: "ctrl+c", help: "exit"},
			{key: "?", help: "help menu"},
		},
	}
}

func (m Model) View() string {
	return "hiiii"
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) SetWH(width int, height int) Model {
	m.Width = width
	m.Height = width
	return m
}
