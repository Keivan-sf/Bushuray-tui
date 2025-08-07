package tunview

import (
	cmds "bushuray-tui/commands"
	"bushuray-tui/utils"

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

func (m Model) View() string {
	dialog := lipgloss.NewStyle().Width(m.Width).MarginTop(2).Padding(0, 2).MaxWidth(m.Width).Align(lipgloss.Center).Render("Core is not running as root, do you want to kill it so you can start again with `sudo bushuray` ?")

	help := lipgloss.NewStyle().Width(m.Width).MarginTop(2).MaxWidth(m.Width).Align(lipgloss.Center).Render(utils.GenHelp([]string{"enter", "esc"}, []string{"YES", "NO"}))

	content := lipgloss.JoinVertical(lipgloss.Top, dialog, help)
	return content
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
	m.Height = width
	return m
}
