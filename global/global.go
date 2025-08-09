package global

import (
	appconfig "bushuray-tui/lib/AppConfig"
	"log"

	"github.com/charmbracelet/lipgloss"
)

var NoColor = lipgloss.NoColor{}
var BgColor = lipgloss.Color("#000001")

func GetBgColor() lipgloss.TerminalColor {
	if appconfig.GetConfig().NoBackground {
		log.Println("giving no color")
		return NoColor
	} else {
		return BgColor
	}
}
