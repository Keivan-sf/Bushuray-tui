package mainmodel

import (
	list "github.com/Keivan-sf/Bushuray-tui/components/List"
	sharedtypes "github.com/Keivan-sf/Bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func applyProfileUpdated(msg sharedtypes.ProfileUpdated, m Model) (tea.Model, tea.Cmd) {
	i, j := findProfile(msg.Profile.GroupId, msg.Profile.Id, m)
	if i != -1 && j != -1 {
		m.Tabs.Children[i].Content.Items[j] = list.ListItem{
			ProfileId:  msg.Profile.Id,
			Name:       msg.Profile.Name,
			Protocol:   convertProtocolForDisplay(msg.Profile.Protocol),
			TestResult: msg.Profile.TestResult,
			Uri:        msg.Profile.Uri,
		}
	}
	return m, nil
}
