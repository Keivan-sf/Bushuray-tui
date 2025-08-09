package mainmodel

import (
	servercmds "bushuray-tui/lib/ServerCommands"
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func applyIsRootAnswer(msg sharedtypes.IsRootAnswer, m Model) (tea.Model, tea.Cmd) {
	if !msg.IsRoot {
		m.ActiveSection = "tunview"
		return m, nil
	}
	servercmds.EnableTun()
	return m, nil
}
