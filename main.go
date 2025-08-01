package main

import (
	"bushuray-tui/components/List"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ConfigList list.Model
	width      int
	height     int
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
		m.ConfigList.Width = msg.Width
		m.ConfigList.Height = msg.Height / 2
	}

	var cmd tea.Cmd
	m.ConfigList, cmd = m.ConfigList.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.ConfigList.View()
}

func initModel() Model {
	return Model{
		ConfigList: list.Model{
			Primary: -1,
			Items: []list.ListItem{
				{Name: "🚀 @SmoothVPN - D", Protocol: "V-LESS"},
				{Name: "[🇨🇦]t.me/ConfigsHub", Protocol: "SHADOW"},
				{Name: "ids here", Protocol: "V-MESS"},
				{Name: "join here @somerandomid", Protocol: "TROJAN"},
				{Name: "50gb config", Protocol: "TROJAN"},
				{Name: "پروکسی دایمی", Protocol: "SOCKS5"},
				{Name: "🚀 @SmoothVPN - D", Protocol: "V-LESS"},
				{Name: "[🇨🇦]t.me/ConfigsHub", Protocol: "SHADOW"},
				{Name: "ids here", Protocol: "V-MESS"},
				{Name: "join here @somerandomid", Protocol: "TROJAN"},
				{Name: "50gb config", Protocol: "TROJAN"},
				{Name: "پروکسی دایمی", Protocol: "SOCKS5"},
				{Name: "🚀 @SmoothVPN - D", Protocol: "V-LESS"},
				{Name: "[🇨🇦]t.me/ConfigsHub", Protocol: "SHADOW"},
				{Name: "ids here", Protocol: "V-MESS"},
				{Name: "join here @somerandomid", Protocol: "TROJAN"},
				{Name: "50gb config", Protocol: "TROJAN"},
				{Name: "پروکسی دایمی", Protocol: "SOCKS5"},
				{Name: "🚀 @SmoothVPN - D", Protocol: "V-LESS"},
				{Name: "[🇨🇦]t.me/ConfigsHub", Protocol: "SHADOW"},
				{Name: "ids here", Protocol: "V-MESS"},
				{Name: "join here @somerandomid", Protocol: "TROJAN"},
				{Name: "50gb config", Protocol: "TROJAN"},
				{Name: "پروکسی دایمی", Protocol: "SOCKS5"},
				{Name: "🚀 @SmoothVPN - D", Protocol: "V-LESS"},
				{Name: "[🇨🇦]t.me/ConfigsHub", Protocol: "SHADOW"},
				{Name: "ids here", Protocol: "V-MESS"},
				{Name: "join here @somerandomid", Protocol: "TROJAN"},
				{Name: "50gb config", Protocol: "TROJAN"},
				{Name: "پروکسی دایمی", Protocol: "SOCKS5"},
				{Name: "🚀 @SmoothVPN - D", Protocol: "V-LESS"},
				{Name: "[🇨🇦]t.me/ConfigsHub", Protocol: "SHADOW"},
				{Name: "ids here", Protocol: "V-MESS"},
				{Name: "join here @somerandomid", Protocol: "TROJAN"},
				{Name: "50gb config", Protocol: "TROJAN"},
				{Name: "پروکسی دایمی", Protocol: "SOCKS5"},
			}},
	}
}

func main() {
	p := tea.NewProgram(initModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}
}
