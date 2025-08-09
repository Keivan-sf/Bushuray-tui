package tunview

import (
	cmds "bushuray-tui/commands"
	"bushuray-tui/components/shared"
	"bushuray-tui/global"
	servercmds "bushuray-tui/lib/ServerCommands"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Width  int
	Height int
}

func InitialModel() Model {
	return Model{}
}

var dialog_style = lipgloss.NewStyle().Padding(0, 2).Align(lipgloss.Center).Background(global.GetBgColor())
var help_style = lipgloss.NewStyle().MarginTop(2).Align(lipgloss.Center).Background(global.GetBgColor())
var help_test = shared.GenHelp([]string{"enter", "esc"}, []string{"YES", "NO"})

func (m Model) View() string {
	bg_style := lipgloss.NewStyle().Background(global.GetBgColor())
	var sudo_bushuray = lipgloss.NewStyle().Foreground(lipgloss.Color("#ca9ee6")).Background(global.GetBgColor()).Render("sudo bushuray")
	var dialog_text = lipgloss.JoinHorizontal(lipgloss.Left, bg_style.Render("Core is not running as root, do you want to kill it so you can start again with "), sudo_bushuray, bg_style.Render(" ?"))
	dialog := dialog_style.Width(m.Width).MaxWidth(m.Width).Background(global.GetBgColor()).Render(dialog_text)
	help := help_style.Width(m.Width).MaxWidth(m.Width).Background(global.GetBgColor()).Render(help_test)
	content := lipgloss.JoinVertical(lipgloss.Top, dialog, help)
	container := lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, content)
	container_with_bg := lipgloss.NewStyle().Background(global.GetBgColor()).Render(container)
	return container_with_bg
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			servercmds.Die()
			return m, tea.Quit
		case "esc":
			return m, cmds.ExitTunView
		}
	}
	return m, nil
}

func (m Model) SetWH(width int, height int) Model {
	m.Width = width
	m.Height = height
	return m
}
