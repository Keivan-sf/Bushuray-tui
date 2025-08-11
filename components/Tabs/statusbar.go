package tabs

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderStatusBar() string {
	bg_width := m.Width
	connection_status := ""
	tun_status := ""

	if m.IsConnected {
		bg_width -= 11
		connection_status = lipgloss.NewStyle().Background(lipgloss.Color("#8839ef")).Foreground(lipgloss.Color("#ffffff")).Width(11).Align(lipgloss.Center).Render("CONNECTED")
	} else {
		bg_width -= 15
		connection_status = lipgloss.NewStyle().Background(lipgloss.Color("#1e1e2e")).Foreground(lipgloss.Color("#ffffff")).Width(15).Align(lipgloss.Center).Render("NOT CONNECTED")
	}

	if m.TunStatus == "connected" {
		bg_width -= 8
		tun_status = lipgloss.NewStyle().Background(lipgloss.Color("#8839ef")).Foreground(lipgloss.Color("#ffffff")).Width(8).Align(lipgloss.Center).Render("TUN ON")
	} else if m.TunStatus == "waiting" {
		bg_width -= 9
		tun_status = lipgloss.NewStyle().Background(lipgloss.Color("#1e1e2e")).Foreground(lipgloss.Color("#ffffff")).Width(9).Align(lipgloss.Center).Render("WAITING")
	} else {
		bg_width -= 9
		tun_status = lipgloss.NewStyle().Background(lipgloss.Color("#1e1e2e")).Foreground(lipgloss.Color("#ffffff")).Width(9).Align(lipgloss.Center).Render("TUN OFF")
	}

	bg := lipgloss.NewStyle().Width(bg_width).Background(lipgloss.Color("#11111b")).Foreground(lipgloss.Color("#585b70")).Render(fmt.Sprintf(" SOCKS5 *:%d | HTTP *:%d", m.SocksPort, m.HttpPort))

	return lipgloss.NewStyle().MarginTop(2).Render(lipgloss.JoinHorizontal(lipgloss.Top, connection_status, bg, tun_status))
}
