package main

import (
	mainmodel "bushuray-tui/components/MainModel"
	connection "bushuray-tui/lib/Connection"
	servercmds "bushuray-tui/lib/ServerCommands"
	servernotifs "bushuray-tui/lib/ServerNotifs"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)
import lumberjack "gopkg.in/natefinch/lumberjack.v2"

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "debug.log",
		MaxSize:    20,
		MaxBackups: 1,
		MaxAge:     0,
		Compress:   false,
	})
	log.SetPrefix("debug: ")
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)

	C := connection.ConnectionHandler{}
	C.Init("127.0.0.1", 4897)

	err := C.GetConnection()
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
