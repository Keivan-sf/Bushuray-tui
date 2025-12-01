package global

import (
	appconfig "github.com/Keivan-sf/Bushuray-tui/lib/AppConfig"

	"github.com/charmbracelet/lipgloss"
)

var NoColor = lipgloss.NoColor{}
var BgColor = lipgloss.Color("#020308")

func GetBgColor() lipgloss.TerminalColor {
	if appconfig.GetConfig().NoBackground {
		return NoColor
	} else {
		return BgColor
	}
}
