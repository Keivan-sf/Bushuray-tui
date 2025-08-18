package list

import (
	servercmds "bushuray-tui/lib/ServerCommands"
	sharedtypes "bushuray-tui/shared_types"
	"github.com/atotto/clipboard"
	"log"
)

func (l *Model) deleteProfileUnderCursor() {
	if l.cursor == l.Primary || len(l.Items) < 1 {
		return
	}
	servercmds.DeleteProfiles([]sharedtypes.ProfileID{{Id: l.Items[l.cursor].ProfileId, GroupId: l.GroupId}})
}

func (l *Model) paste() {
	str, err := clipboard.ReadAll()
	if err != nil {
		log.Println("There was an error readin from clipboard", err)
		return
	}
	servercmds.AddProfiles(str, l.GroupId)
}

func (l *Model) copyProfileUnderCursor() {
	if len(l.Items) < 1 {
		return
	}
	uri := l.Items[l.cursor].Uri
	err := clipboard.WriteAll(uri)
	if err != nil {
		log.Println("There was an error writing to clipboard", err)
		return
	}
}

func (l *Model) testProfile() {
	if len(l.Items) < 1 {
		return
	}
	l.Items[l.cursor].TestResult = -2
	servercmds.Test(l.GroupId, l.Items[l.cursor].ProfileId)
}

func (l *Model) testGroup() {
	for i, item := range l.Items {
		l.Items[i].TestResult = -2
		servercmds.Test(l.GroupId, item.ProfileId)
	}
}

func (l *Model) connectToProfile() {
	if len(l.Items) < 1 {
		return
	}
	if l.Primary == l.cursor {
		servercmds.Disconnect()
	} else {
		l.Primary = l.cursor
		servercmds.Connect(l.GroupId, l.Items[l.Primary].ProfileId)
	}
}
