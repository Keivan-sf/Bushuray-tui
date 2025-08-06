package cmds

import (
	t "bushuray-tui/shared_types"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func SubmitGroup(title string, url string) tea.Cmd {
	return func() tea.Msg {
		// send data to core
		log.Println(title, url)
		return t.AddGroupExit{}
	}
}

func ExitAddGroupView() tea.Msg {
	return t.AddGroupExit{}
}

func EnterAddGroupView() tea.Msg {
	return t.AddGroupEnter{}
}
