package tunview

import (
	cmds "bushuray-tui/commands"
	"bushuray-tui/utils"
	"fmt"

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

var sudo_bushuray = lipgloss.NewStyle().Foreground(lipgloss.Color("#ca9ee6")).Render("sudo bushuray")
var dialog_text = fmt.Sprintf("Core is not running as root, do you want to kill it so you can start again with %s ?", sudo_bushuray)
var dialog_style = lipgloss.NewStyle().Padding(0, 2).Align(lipgloss.Center)
var help_style = lipgloss.NewStyle().MarginTop(2).Align(lipgloss.Center)
var help_test = utils.GenHelp([]string{"enter", "esc"}, []string{"YES", "NO"})

func (m Model) View() string {
	dialog := dialog_style.Width(m.Width).MaxWidth(m.Width).Render(dialog_text)
	help := help_style.Width(m.Width).MaxWidth(m.Width).Render(help_test)
	content := lipgloss.JoinVertical(lipgloss.Top, dialog, help)
	container := lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, content)
	return container
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, tea.Batch(cmds.KillCore, tea.Quit)
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
