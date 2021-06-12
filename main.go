package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

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

	// Console
	fmt.Println("Yo")
	if Xonsole.Cmd.Start() != nil {
		log.Fatal("Cannot start server")
	}

	go func() {
		for {
			if x, err := Xonsole.ReadLine(); err != io.EOF {
				fmt.Println(x)
			}
		}
	}()

	for {
	}

}
