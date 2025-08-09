package mainmodel

import (
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func applyGroupDeleted(msg sharedtypes.GroupDeleted, m Model) (tea.Model, tea.Cmd) {
	tid := findGroupTab(msg.Id, m)
	if tid != -1 {
		m.Tabs.DeleteTab(tid)
	}
	return m, nil
}
