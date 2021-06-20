package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

var Xonsole *Console

var socketConns *map[string]socketio.Conn = &map[string]socketio.Conn{}

var allowOriginFunc = func(r *http.Request) bool {
	// Cors (*)
	return true
}

// TODO : Make a log parser and do something about restarts and server states in frontend also if mood permits organize this code
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
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	go handlesocketConns(socketConns)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		s.Emit("reply", "Successfully Connected 😀")
		s.Emit("reply", "Type help for more info")
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
		} else if msg == "stop" || msg == "/stop" {
			Xonsole.ExecCommand("stop")
			// TODO : maybe check for eof then decrement waitgroup and terminate the loop and exit gracefully
			time.Sleep(5 * time.Second)
			Xonsole.Cmd.Process.Kill()
			s.Emit("reply", "Shutting down server disconnected contact admin if you need to restart")
			s.Emit("reply", "PS: In future you will be able to start this server from browser only")
			os.Exit(0)
		} else {
			Xonsole.ExecCommand(msg)

		}
	})

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
	app.Router.Handle("/ws/", server)

	for {
	}
}

const (
	Starting          = iota // Preparing
	Stopping                 // Saving
	Restarting               // Custom
	Online                   // help
	Offline                  // Xonsole is nil also give this a littlebit of a thought can be EOF but meh that not good
	FailedPortBinding        // FAILED TO BIND TO PORT
	UnknownError             // Think of it
)

func parselogs(x string) {
	out := strings.Split(x, ":")
	toParse := out[len(out)-1]
	fmt.Println(toParse)
	if strings.Contains(toParse, "Done") {
		fmt.Println("Satrted server successfully!!")
	} else if strings.Contains(toParse, "EULA") {
		fmt.Println("You havent agreed to eula first agree to eula. Stopping server now")
	} else if strings.Contains(toParse, "FAILED TO BIND TO PORT") {
		fmt.Println("Perhaps another server instance is runnig stop that first")
	}
}

func handlesocketConns(sockets *map[string]socketio.Conn) {
	for {
		if x, err := Xonsole.ReadLine(); err != io.EOF {
			// fmt.Println(x, " ", len(*sockets))
			parselogs(x)
			for _, value := range *sockets {
				value.Emit("reply", x)
			}
		} else {
			fmt.Println(err)
		}
	}
}
