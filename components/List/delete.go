package list

import (
	servercmds "bushuray-tui/lib/ServerCommands"
	sharedtypes "bushuray-tui/shared_types"
)

func (l Model) deleteProfileUnderCursor() {
	if l.cursor == l.Primary || l.cursor >= len(l.Items) {
		return
	}
	servercmds.DeleteProfiles([]sharedtypes.ProfileID{{Id: l.Items[l.cursor].ProfileId, GroupId: l.GroupId}})
}
