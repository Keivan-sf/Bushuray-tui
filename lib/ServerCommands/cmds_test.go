package servercmds

import (
	"testing"

	sharedtypes "github.com/Keivan-sf/Bushuray-tui/shared_types"
)

func TestCreateStopTestsCommand(t *testing.T) {
	got := string(CreateJsonCommand("stop-tests", sharedtypes.StopTestsData{}))
	want := `{"msg":"stop-tests","data":{}}`

	if got != want {
		t.Fatalf("CreateJsonCommand() = %s, want %s", got, want)
	}
}
