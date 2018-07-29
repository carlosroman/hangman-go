package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"hangman/server/rest"
	"hangman/services"
	"hangman/services/wordstore"
	"hangman/utils"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	utils.EnableServerMode()
	var port int
	var csvPath string
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
		cli.StringFlag{
			Name:        "csv, c",
			Usage:       "File path to CSV file",
			EnvVar:      "CSV_PATH",
			Destination: &csvPath,
		},
	}

	app.Action = func(c *cli.Context) error {
		r := mux.NewRouter()
		ws, _ := wordstore.NewInMemoryStoreFromCSV(csvPath) // todo: deal with error
		gs := rest.NewGameServer(r, services.NewGameService(ws))
		gs.InitialiseHandlers()
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
