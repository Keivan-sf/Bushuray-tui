package main

import (
	mainmodel "bushuray-tui/components/MainModel"
	connection "bushuray-tui/lib/Connection"
	servercmds "bushuray-tui/lib/ServerCommands"
	servernotifs "bushuray-tui/lib/ServerNotifs"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("error with log to file", err)
	}
	defer f.Close()
	C := connection.ConnectionHandler{}
	C.Init("127.0.0.1", 4897)
	err = C.GetConnection()
	if err != nil {
		fmt.Println("error connecting", err)
		return
	} else {
		fmt.Println("connection established")
	}

	zone.NewGlobal()
	p := tea.NewProgram(mainmodel.InitModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())

	go C.HandleConnection(p)
	servercmds.Init(&C)
	servercmds.GetApplicationState()
	servernotifs.Init(p)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}
}
