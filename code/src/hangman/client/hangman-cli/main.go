package main

import (
	"github.com/rivo/tview"
	"github.com/urfave/cli"
	hangman "hangman/client/hangman-cli/app"
	"os"
)

func main() {

	var baseUrl string
	app := cli.NewApp()
	app.Name = "Hangman"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "baseUrl",
			Value:       "http://localhost:8080",
			Usage:       "Set the base URL to connect to the server",
			EnvVar:      "BASE_URL",
			Destination: &baseUrl,
		},
	}

	app.Action = func(c *cli.Context) error {
		a := hangman.NewApp(&baseUrl)

		pages := tview.NewPages()

		nextSlide := func(p string) {
			pages.SwitchToPage(p)
		}

		pages.AddPage(
			"welcome",
			hangman.Welcome(nextSlide, a.StopApp),
			true,
			true,
		)

		pages.AddPage(
			"start",
			hangman.Start(nextSlide, a.StopApp, a.NewGame),
			true,
			false,
		)

		info := tview.NewTextView().
			SetDynamicColors(true).
			SetRegions(true).
			SetWrap(false)

		// App layout
		layout := tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(pages, 0, 1, true).
			AddItem(info, 1, 1, false)

		return a.SetRoot(layout, true).Run()
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
