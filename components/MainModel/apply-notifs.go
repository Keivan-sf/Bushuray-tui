package mainmodel

import (
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func HandleServerNotifs(msg sharedtypes.ServerNotification, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case sharedtypes.ApplicationState:
		return applyApplicationState(msg, m)
	}
	return m, nil
}
