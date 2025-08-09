package mainmodel

import (
	list "bushuray-tui/components/List"
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func applyProfilesAdded(msg sharedtypes.ProfilesAdded, m Model) (tea.Model, tea.Cmd) {
	for _, profile := range msg.Profiles {
		tid := findGroupTab(profile.GroupId, m)
		if tid == -1 {
			continue
		}
		m.Tabs.Children[tid].Content.Items = append(m.Tabs.Children[tid].Content.Items, list.ListItem{
			ProfileId:  profile.Id,
			Name:       profile.Name,
			Protocol:   profile.Protocol,
			TestResult: profile.TestResult,
			Uri:        profile.Uri,
		})
	}
	return m, nil
}
