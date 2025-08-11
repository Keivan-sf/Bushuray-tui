package mainmodel

import (
	addgroup "bushuray-tui/components/AddGroup"
	helpview "bushuray-tui/components/Help"
	list "bushuray-tui/components/List"
	tabs "bushuray-tui/components/Tabs"
	tunview "bushuray-tui/components/Tun"
	appconfig "bushuray-tui/lib/AppConfig"
	sharedtypes "bushuray-tui/shared_types"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

type Model struct {
	Width         int
	Height        int
	Tabs          tabs.Model
	AddGroup      addgroup.Model
	Tun           tunview.Model
	Help          helpview.Model
	ActiveSection string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case sharedtypes.ServerNotification:
		return HandleServerNotifs(msg, m)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Tabs = m.Tabs.SetWH(msg.Width, msg.Height)
		m.AddGroup = m.AddGroup.SetWH(msg.Width, msg.Height)
		m.Tun = m.Tun.SetWH(msg.Width, msg.Height)
		m.Help = m.Help.SetWH(msg.Width, msg.Height)
		return m, nil

	case sharedtypes.AddGroupExit:
		m.ActiveSection = "tabs"
		return m, nil

	case sharedtypes.AddGroupEnter:
		m.ActiveSection = "add-group"
		return m, nil

	case sharedtypes.TunViewEnter:
		m.ActiveSection = "tunview"
		return m, nil

	case sharedtypes.TunViewExit:
		m.ActiveSection = "tabs"
		return m, nil

	case sharedtypes.HelpViewEnter:
		m.ActiveSection = "helpview"
		return m, nil

	case sharedtypes.HelpViewExit:
		m.ActiveSection = "tabs"
		return m, nil
	}

	if m.ActiveSection == "tunview" {
		var cmd tea.Cmd
		m.Tun, cmd = m.Tun.Update(msg)
		return m, cmd
	}

	if m.ActiveSection == "add-group" {
		var cmd tea.Cmd
		m.AddGroup, cmd = m.AddGroup.Update(msg)
		return m, cmd
	}

	if m.ActiveSection == "helpview" {
		var cmd tea.Cmd
		m.Help, cmd = m.Help.Update(msg)
		return m, cmd
	}

	if m.ActiveSection == "tabs" {
		var cmd tea.Cmd
		m.Tabs, cmd = m.Tabs.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	if m.ActiveSection == "helpview" {
		return m.Help.View()
	}
	if m.ActiveSection == "add-group" {
		return m.AddGroup.View()
	}
	if m.ActiveSection == "tunview" {
		return m.Tun.View()
	}
	return zone.Scan(m.Tabs.View())
}

func InitModel() Model {
	config := appconfig.GetConfig()
	return Model{
		ActiveSection: "tabs",
		Tun:           tunview.InitialModel(),
		AddGroup:      addgroup.InitialModel(),
		Help:          helpview.InitialModel(),
		Tabs: tabs.Model{
			Id:          zone.NewPrefix(),
			TunStatus:   "disconnected",
			IsConnected: true,
			SocksPort:   config.SocksPort,
			HttpPort:    config.HttpPort,
			Children: []tabs.TabView{
				{
					Content: list.Model{
						Id:      zone.NewPrefix(),
						Primary: -1,
						Items:   []list.ListItem{},
					},
					Title: "Default group",
				},
			},
		},
	}
}
