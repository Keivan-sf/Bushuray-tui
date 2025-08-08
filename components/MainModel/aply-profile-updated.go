package mainmodel

import (
	list "bushuray-tui/components/List"
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func applyProfileUpdated(msg sharedtypes.ProfileUpdated, m Model) (tea.Model, tea.Cmd) {
	for i, g := range m.Tabs.Children {
		if g.Content.GroupId == msg.Profile.GroupId {
			for j, p := range g.Content.Items {
				if p.ProfileId == msg.Profile.Id {
					m.Tabs.Children[i].Content.Items[j] = list.ListItem{
						ProfileId:  msg.Profile.Id,
						Name:       msg.Profile.Name,
						Protocol:   convertProtocolForDisplay(msg.Profile.Protocol),
						TestResult: msg.Profile.TestResult,
						Uri:        msg.Profile.Uri,
					}
					break
				}
			}
			break
		}
	}
	return m, nil
}
