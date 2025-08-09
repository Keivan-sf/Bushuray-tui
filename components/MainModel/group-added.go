package mainmodel

import (
	list "bushuray-tui/components/List"
	tabs "bushuray-tui/components/Tabs"
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func applyGroupAdded(msg sharedtypes.GroupAdded, m Model) (tea.Model, tea.Cmd) {
	tabview := tabs.TabView{
		Title: msg.Name,
		Content: list.Model{
			Id:      zone.NewPrefix(),
			GroupId: msg.Id,
			Primary: -1,
			Items:   []list.ListItem{},
		},
	}

	tid := findGroupTab(msg.Id, m)
	if tid != -1 {
		m.Tabs.Children[tid] = tabview
	} else {
		m.Tabs.Children = append(m.Tabs.Children, tabview)
	}
	m.Tabs = m.Tabs.SetWH(m.Tabs.Width, m.Tabs.Height)
	return m, nil
}
