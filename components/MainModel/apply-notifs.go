package mainmodel

import (
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
)

func HandleServerNotifs(msg sharedtypes.ServerNotification, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case sharedtypes.Warning:
		return applyWarning(msg, m)
	case sharedtypes.ApplicationState:
		return applyApplicationState(msg, m)
	case sharedtypes.ProfileUpdated:
		return applyProfileUpdated(msg, m)
	case sharedtypes.ProxyStatus:
		return applyStatusChanged(msg, m)
	case sharedtypes.ProfilesAdded:
		return applyProfilesAdded(msg, m)
	case sharedtypes.ProfilesDeleted:
		return applyProfilesDeleted(msg, m)
	case sharedtypes.GroupAdded:
		return applyGroupAdded(msg, m)
	case sharedtypes.GroupDeleted:
		return applyGroupDeleted(msg, m)
	case sharedtypes.SubscriptionUpdated:
		return applySubscriptionUpdated(msg, m)
	case sharedtypes.IsRootAnswer:
		return applyIsRootAnswer(msg, m)
	case sharedtypes.TunStatus:
		return applyTunStatusChanged(msg, m)
	}
	return m, nil
}

func findProfile(gid int, id int, m Model) (tabid int, index int) {
	for i, g := range m.Tabs.Children {
		if g.Content.GroupId == gid {
			for j, p := range g.Content.Items {
				if p.ProfileId == id {
					return i, j
				}
			}
		}
	}
	return -1, -1
}

func findGroupTab(gid int, m Model) (tabid int) {
	for i, g := range m.Tabs.Children {
		if g.Content.GroupId == gid {
			return i
		}
	}
	return -1
}
