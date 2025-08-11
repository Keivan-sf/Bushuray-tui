package mainmodel

import (
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func applyTunStatusChanged(msg sharedtypes.TunStatus, m Model) (tea.Model, tea.Cmd) {
	if msg.IsEnabled {
		m.Tabs.TunStatus = "connected"
	} else {
		m.Tabs.TunStatus = "disconnected"
	}
	return m, nil
}
