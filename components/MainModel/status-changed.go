package mainmodel

import (
	sharedtypes "github.com/Keivan-sf/Bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func applyStatusChanged(msg sharedtypes.ProxyStatus, m Model) (tea.Model, tea.Cmd) {
	m.Tabs.IsConnected = msg.Connection == "connected"
	if m.Tabs.IsConnected {
		tid, id := findProfile(msg.Profile.GroupId, msg.Profile.Id, m)
		if tid != -1 && id != -1 {
			m.Tabs.Children[tid].Content.Primary = id
		}
	} else {
		for i := range m.Tabs.Children {
			m.Tabs.Children[i].Content.Primary = -1
		}
	}
	return m, nil
}
