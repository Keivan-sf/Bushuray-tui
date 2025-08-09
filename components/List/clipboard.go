package list

import (
	servercmds "bushuray-tui/lib/ServerCommands"
	"log"

	"github.com/atotto/clipboard"
)

func (l Model) paste() {
	str, err := clipboard.ReadAll()
	if err != nil {
		log.Println("There was an error readin from clipboard", err)
		return
	}
	servercmds.AddProfiles(str, l.GroupId)
}

func (l Model) copyUri(uri string) {
	err := clipboard.WriteAll(uri)
	if err != nil {
		log.Println("There was an error writing to clipboard", err)
		return
	}
}
