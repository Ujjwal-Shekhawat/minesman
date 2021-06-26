package main

import (
	"fmt"
	"log"

	"github.com/Ujjwal-Shekhawat/minesman/console"
	server "github.com/Ujjwal-Shekhawat/minesman/http_socket_server"
)

var Xonsole = console.Xonsole

// TODO : Make a log parser and do something about restarts and server states in frontend also if mood permits organize this code
func main() {
	// Console
	(*Xonsole)["exp"] = console.InitConsole()
	fmt.Println("Yo")
	for sName := range *Xonsole {
		if (*Xonsole)[sName].Cmd.Start() != nil {
			log.Fatal("Cannot start server : " + sName)
		}
	}

	server.StartServer(Xonsole)
}
