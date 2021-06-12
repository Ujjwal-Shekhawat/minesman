package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var Xonsole = InitConsole()

func main() {
	// Server
	port := os.Getenv("PORT")
	app := App{}
	app.initRoutes()
	if port == "" {
		port = "8080"
	}
	go app.run(port)
	app.Router.HandleFunc("/ws", wsConnection).Methods("GET")

	// Socker Server

	// Console
	fmt.Println("Yo")
	if Xonsole.Cmd.Start() != nil {
		log.Fatal("Cannot start server")
	}

	/* go func() {
		for {
			if x, err := Xonsole.ReadLine(); err != io.EOF {
				fmt.Println(x)
			}
		}
	}() */

	for {
	}

}

func reader(reader *websocket.Conn) {
	for {
		mess, p, err := reader.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		// Client message here
		fmt.Println(string(p))
		if string(p) != "" {
			Xonsole.ExecCommand(string(p))
		}

		if err := reader.WriteMessage(mess, p); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func wsConnection(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ M string }{"Socket conn upgrade failed"})
	}
	log.Println("A new client connected")

	go func(ws *websocket.Conn) {
		for {
			if x, err := Xonsole.ReadLine(); err != io.EOF {
				fmt.Println(x)
				err = ws.WriteMessage(1, []byte(x))
				if err != nil {
					log.Println("unable to write message")
					break
				}
			}
		}
		return
	}(ws)
	err = ws.WriteMessage(1, []byte("Hello from Server"))
	if err != nil {
		log.Fatal("unable to write message")
	}

	reader(ws)
}
