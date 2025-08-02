package main

import (
	"bushuray-tui/components/List"
	tabs "bushuray-tui/components/Tabs"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

type Model struct {
	width  int
	height int
	tabs   tabs.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.tabs = m.tabs.SetWH(msg.Width, msg.Height/2)
		return m, nil
	}

	var cmd tea.Cmd
	m.tabs, cmd = m.tabs.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return zone.Scan(m.tabs.View())
}

func initModel() Model {
	return Model{
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
						Items:   dummy_items,
					},
					Title: "Group 0",
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
