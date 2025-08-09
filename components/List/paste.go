package list

import (
	servercmds "bushuray-tui/lib/ServerCommands"
	"log"

	"github.com/atotto/clipboard"
)

func (l Model) paste() {
	str, err := clipboard.ReadAll()
	if err != nil {
		log.Println("There was an error writing to clipboard", err)
	}
	servercmds.AddProfiles(str, l.GroupId)
}
