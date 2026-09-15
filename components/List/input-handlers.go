package list

import (
	"log"

	appconfig "github.com/Keivan-sf/Bushuray-tui/lib/AppConfig"
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
	l.testQueue = []int{}
	for i, item := range l.Items {
		l.Items[i].TestResult = 0
		l.testQueue = append(l.testQueue, item.ProfileId)
	}
	l.groupTesting = true
	l.ContinueGroupTest()
}

func (l *Model) ContinueGroupTest() {
	if !l.groupTesting {
		return
	}
	config := appconfig.GetConfig()
	testing := 0
	working := 0
	for _, item := range l.Items {
		if item.TestResult == -2 {
			testing++
		} else if item.TestResult > 0 {
			working++
		}
	}
	if config.StopTestingAfter > 0 && working >= config.StopTestingAfter {
		l.testQueue = nil
	}
	for len(l.testQueue) > 0 && (config.StopTestingAfter <= 0 || testing < test_concurrency) {
		id := l.testQueue[0]
		l.testQueue = l.testQueue[1:]
		for i, item := range l.Items {
			if item.ProfileId == id {
				l.Items[i].TestResult = -2
				servercmds.Test(l.GroupId, id)
				testing++
				break
			}
		}
	}
	if testing == 0 && len(l.testQueue) == 0 {
		l.groupTesting = false
		if config.RemoveFailedProfiles {
			l.RemoveFailedProfiles()
		}
	}
}

func (l *Model) RemoveFailedProfiles() {
	failed := []sharedtypes.ProfileID{}
	for i, item := range l.Items {
		if item.TestResult == -1 && i != l.Primary {
			failed = append(failed, sharedtypes.ProfileID{Id: item.ProfileId, GroupId: l.GroupId})
		}
	}
	if len(failed) > 0 {
		servercmds.DeleteProfiles(failed)
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
