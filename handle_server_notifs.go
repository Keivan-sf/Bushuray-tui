package main

import (
	sharedtypes "bushuray-tui/shared_types"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func HandleServerNotifs(msg sharedtypes.ServerNotification, m tea.Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case sharedtypes.ApplicationState:
		log.Printf("received application state notification:", msg)
	}
	return m, nil
}
