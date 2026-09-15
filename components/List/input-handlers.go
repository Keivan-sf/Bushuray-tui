package list

import (
	"log"

	notif_publisher "github.com/Keivan-sf/Bushuray-tui/lib/NotifPublisher"
	servercmds "github.com/Keivan-sf/Bushuray-tui/lib/ServerCommands"
	sharedtypes "github.com/Keivan-sf/Bushuray-tui/shared_types"
	"github.com/atotto/clipboard"
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
		log.Println("There was a problem reading from clipboard, entering paste profile view mode", err)
		go func() { notif_publisher.EnterPasteProfileMode(sharedtypes.PasteProfileViewEnter{}) }()
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
