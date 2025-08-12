package mainmodel

import (
	list "bushuray-tui/components/List"
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func applySubscriptionUpdated(msg sharedtypes.SubscriptionUpdated, m Model) (tea.Model, tea.Cmd) {
	tid := findGroupTab(msg.GroupId, m)

	if tid == -1 {
		return m, nil
	}

	items := []list.ListItem{}

	for _, profile := range msg.Profiles {
		child := list.ListItem{
			Name:       profile.Name,
			ProfileId:  profile.Id,
			Protocol:   convertProtocolForDisplay(profile.Protocol),
			TestResult: profile.TestResult,
			Uri:        profile.Uri,
		}
		items = append(items, child)
	}

	m.Tabs.Children[tid].Content.Items = items
	if m.Tabs.Children[tid].Content.Primary != -1 {
		m.Tabs.Children[tid].Content.Primary = 0
	}
	m.Tabs.Children[tid].Content.ResetCursor()
	return m, nil
}
