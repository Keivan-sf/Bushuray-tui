package mainmodel

import (
	list "bushuray-tui/components/List"
	tabs "bushuray-tui/components/Tabs"
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func applyApplicationState(msg sharedtypes.ApplicationState, m Model) (tea.Model, tea.Cmd) {
	connected_profile_id := -1
	connected_profile_gid := -1
	if msg.ConnectionStatus.Connection == "connected" {
		connected_profile_id = msg.ConnectionStatus.Profile.Id
		connected_profile_gid = msg.ConnectionStatus.Profile.GroupId
	}

	m.Tabs.IsConnected = msg.ConnectionStatus.Connection == "connected"
	m.Tabs.IsTunEnabled = msg.TunStatus
	var views []tabs.TabView
	for _, group := range msg.Groups {
		var tabview tabs.TabView
		tabview.Content = list.Model{}
		tabview.Content.Primary = -1
		tabview.Content.Items = []list.ListItem{}
		tabview.Content.Id = zone.NewPrefix()
		tabview.Content.GroupId = group.Group.Id
		tabview.Title = group.Group.Name
		for i, profile := range group.Profiles {
			if connected_profile_gid == profile.GroupId && connected_profile_id == profile.Id {
				tabview.Content.Primary = i
			}
			child := list.ListItem{
				Name:       profile.Name,
				ProfileId:  profile.Id,
				Protocol:   convertProtocolForDisplay(profile.Protocol),
				TestResult: profile.TestResult,
			}
			tabview.Content.Items = append(tabview.Content.Items, child)
		}
		views = append(views, tabview)
	}
	m.Tabs.Children = views
	m.Tabs.ActiveTap = 0
	m.Tabs = m.Tabs.SetWH(m.Tabs.Width, m.Tabs.Height)

	return m, nil
}

func convertProtocolForDisplay(name string) string {
	switch name {
	case "vless":
		return "V-LESS"
	case "vmess":
		return "V-MESS"
	case "socks":
		return "SOCKS5"
	case "shadowsocks":
		return "SHADOW"
	case "trojan":
		return "TROJAN"
	}
	return name
}
