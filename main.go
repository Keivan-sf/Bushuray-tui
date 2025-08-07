package main

import (
	addgroup "bushuray-tui/components/AddGroup"
	"bushuray-tui/components/List"
	tabs "bushuray-tui/components/Tabs"
	tunview "bushuray-tui/components/Tun"
	sharedtypes "bushuray-tui/shared_types"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

type Model struct {
	width          int
	height         int
	tabs           tabs.Model
	add_group      addgroup.Model
	tun            tunview.Model
	active_section string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.tabs = m.tabs.SetWH(msg.Width, msg.Height/2)
		m.add_group = m.add_group.SetWH(msg.Width, msg.Height)
		m.tun = m.tun.SetWH(msg.Width, msg.Height)
		return m, nil

	case sharedtypes.AddGroupExit:
		m.active_section = "tabs"
		return m, nil

	case sharedtypes.AddGroupEnter:
		m.active_section = "add-group"
		return m, nil
	}

	if m.active_section == "tunview" {
		return m, nil
	}

	if m.active_section == "add-group" {
		var cmd tea.Cmd
		m.add_group, cmd = m.add_group.Update(msg)
		return m, cmd
	}

	if m.active_section == "tabs" {
		var cmd tea.Cmd
		m.tabs, cmd = m.tabs.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	if m.active_section == "add-group" {
		return m.add_group.View()
	}
	if m.active_section == "tunview" {
		return m.tun.View()
	}
	return zone.Scan(m.tabs.View())
}

func initModel() Model {
	return Model{
		active_section: "tunview",
		tun:            tunview.InitialModel(),
		add_group:      addgroup.InitialModel(),
		tabs: tabs.Model{
			Id: zone.NewPrefix(),
			Children: []tabs.TabView{
				{
					Content: list.Model{
						Id:      zone.NewPrefix(),
						Primary: -1,
						Items:   dummy_items,
					},
					Title: "Default group",
				},
				{
					Content: list.Model{
						Id:      zone.NewPrefix(),
						Primary: -1,
						Items:   dummyItemsWithPrefix("G1"),
					},
					Title: "mt Group 1",
				},
				{
					Content: list.Model{
						Id:      zone.NewPrefix(),
						Primary: -1,
						Items:   dummyItemsWithPrefix("G2"),
					},
					Title: "random name",
				},
				{
					Content: list.Model{
						Id:      zone.NewPrefix(),
						Primary: -1,
						Items:   dummyItemsWithPrefix("G3"),
					},
					Title: "Group 3",
				},
				{
					Content: list.Model{
						Id:      zone.NewPrefix(),
						Primary: -1,
						Items:   dummyItemsWithPrefix("g4long"),
					},
					Title: "Group 4 very long name",
				},
				{
					Content: list.Model{
						Id:      zone.NewPrefix(),
						Primary: -1,
						Items:   dummyItemsWithPrefix("A-"),
					},
					Title: "A Group 3-",
				},
				{
					Content: list.Model{
						Id:      zone.NewPrefix(),
						Primary: -1,
						Items:   dummyItemsWithPrefix("BBB"),
					},
					Title: "BBB Group 2",
				},
			},
		},
	}
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("error with log to file", err)
	}
	defer f.Close()

	zone.NewGlobal()
	p := tea.NewProgram(initModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}
}
