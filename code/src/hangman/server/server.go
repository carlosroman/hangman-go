package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"os"
	"time"
	"hangman/services"
	"hangman/server/handlers"
)

func main() {
	var port int
	r := mux.NewRouter()
	gs := handlers.NewGameServer(r, services.NewGameService())
	gs.InitialiseHandlers()
	app := cli.NewApp()
	app.Name = "Hangman server"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "port, p",
			Value:       8080,
			Usage:       "Set the port of the server",
			EnvVar:      "PORT",
			Destination: &port,
		},
	}

	app.Action = func(c *cli.Context) error {
		srv := &http.Server{
			Handler: r,
			Addr:    fmt.Sprintf("127.0.0.1:%d", port),
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		log.Print("Starting server...")
		log.Fatal(srv.ListenAndServe())
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
