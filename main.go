package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

var Xonsole *Console

var socketConns *map[string]socketio.Conn = &map[string]socketio.Conn{}

var allowOriginFunc = func(r *http.Request) bool {
	// Cors (*)
	return true
}

func main() {
	// Console
	Xonsole = InitConsole()
	fmt.Println("Yo")
	if Xonsole.Cmd.Start() != nil {
		log.Fatal("Cannot start server")
	}

	serveAll()
}

func serveAll() {
	// Server
	port := os.Getenv("PORT")
	app := App{}
	app.initRoutes()
	if port == "" {
		port = "8080"
	}
	go app.run(port)

	// Socker Server
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	go handlesocketConns(socketConns)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		// Change this to be like rff or something else
		/* go func() {
			for {
				if x, err := Xonsole.ReadLine(); err != io.EOF {
					fmt.Println(x)
					s.Emit("reply", x)
				} else {
				}
			}
		}() */
		(*socketConns)[s.ID()] = s
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "Executing command :  "+msg)
		if msg == "restart" {
			Xonsole.ExecCommand("stop")
			time.Sleep(5 * time.Second)
			Xonsole.Cmd.Process.Kill()
			// Xonsole = nil Dont know but this line causes memory error
			Xonsole = InitConsole()
			Xonsole.Cmd.Start()
		} else {
			Xonsole.ExecCommand(msg)
		}
	})

	/* server.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	}) */

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		delete(*socketConns, s.ID())
		log.Println("Socket conn closed by client" + s.ID())
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()
	app.Router.Handle("/socket.io/", server)

	for {
	}
}

func handlesocketConns(sockets *map[string]socketio.Conn) {
	for {
		if x, err := Xonsole.ReadLine(); err != io.EOF {
			fmt.Println(x, " ", len(*sockets))
			for _, value := range *sockets {
				value.Emit("reply", x)
			}
		} else {
		}
	}
}
