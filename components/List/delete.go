package list

import (
	servercmds "bushuray-tui/lib/ServerCommands"
	sharedtypes "bushuray-tui/shared_types"
)

func (l Model) deleteProfile(gid int, id int) {
	servercmds.DeleteProfiles([]sharedtypes.ProfileID{{Id: id, GroupId: gid}})
}
