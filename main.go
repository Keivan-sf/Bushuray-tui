package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mainmodel "github.com/Keivan-sf/Bushuray-tui/components/MainModel"
	appconfig "github.com/Keivan-sf/Bushuray-tui/lib/AppConfig"
	connection "github.com/Keivan-sf/Bushuray-tui/lib/Connection"
	notif_publisher "github.com/Keivan-sf/Bushuray-tui/lib/NotifPublisher"
	servercmds "github.com/Keivan-sf/Bushuray-tui/lib/ServerCommands"
	sharedtypes "github.com/Keivan-sf/Bushuray-tui/shared_types"
	"github.com/Keivan-sf/Bushuray-tui/utils"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

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

	appconfig.LoadConfig()
	port := appconfig.GetConfig().CoreTCPPort

	if len(os.Args) > 1 && os.Args[1] == "die" {
		die(port)
		return
	}

	C := connection.ConnectionHandler{}
	C.Init("127.0.0.1", port)

	err := C.GetConnection()
	if err != nil {
		fmt.Println("core was not found at", port, "trying to spawn")
		err := utils.SpawnBushurayCore()
		if err != nil {
			fmt.Println("failed to spawn core:", err)
			return
		}
		time.Sleep(1000 * time.Millisecond)
		err = C.GetConnection()
		if err != nil {
			fmt.Println("failed to connect to core:", err)
			return
		}
	} else {
		fmt.Println("connection established")
	}

	zone.NewGlobal()
	p := tea.NewProgram(mainmodel.InitModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())

	go C.HandleConnection(p)
	servercmds.Init(&C)
	servercmds.GetApplicationState()
	notif_publisher.Init(p)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}
}

func die(port int) {
	C := connection.ConnectionHandler{}
	C.Init("127.0.0.1", port)

	if err := C.GetConnection(); err != nil {
		fmt.Println("Bushuray is not running")
		return
	}
	defer C.Close()

	if err := C.Send(servercmds.CreateJsonCommand("die", sharedtypes.DieData{})); err != nil {
		fmt.Println("failed to stop Bushuray:", err)
		return
	}

	if err := C.WaitForClose(); err != nil {
		fmt.Println("failed while waiting for Bushuray to stop:", err)
		return
	}

	fmt.Println("Bushuray has stopped successfully")
}
