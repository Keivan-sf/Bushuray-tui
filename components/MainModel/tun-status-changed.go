package mainmodel

import (
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func applyTunStatusChanged(msg sharedtypes.TunStatus, m Model) (tea.Model, tea.Cmd) {
	m.Tabs.IsTunEnabled = msg.IsEnabled
	return m, nil
}
