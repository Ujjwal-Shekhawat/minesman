package main

import (
	"fmt"
	"log"

	"github.com/Ujjwal-Shekhawat/minesman/console"
	server "github.com/Ujjwal-Shekhawat/minesman/http_socket_server"
)

var Xonsole *map[string](*console.Console) = &map[string](*console.Console){}

// TODO : Make a log parser and do something about restarts and server states in frontend also if mood permits organize this code
func main() {
	// Console
	(*Xonsole)["exp"] = console.InitConsole()
	fmt.Println("Yo")
	if (*Xonsole)["exp"].Cmd.Start() != nil {
		log.Fatal("Cannot start server")
	}

	server.StartServer(Xonsole)
}
