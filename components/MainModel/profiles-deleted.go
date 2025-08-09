package mainmodel

import (
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func applyProfilesDeleted(msg sharedtypes.ProfilesDeleted, m Model) (tea.Model, tea.Cmd) {
	for _, profile := range msg.DeletedProfiles {
		tid, idx := findProfile(profile.GroupId, profile.Id, m)
		if tid == -1 || idx == -1 {
			continue
		}
		items := m.Tabs.Children[tid].Content.Items
		m.Tabs.Children[tid].Content.Items = append(items[:idx], items[idx+1:]...)
		// cursor might be on the last item when the profile gets deleted
		m.Tabs.Children[tid].Content.ResolveInvalidCursor()
	}
	return m, nil
}
