package main

import (
	"github.com/rivo/tview"
	"hangman/client/hangman-cli/app"
)

func main() {
	a := app.NewApp()

	pages := tview.NewPages()

	nextSlide := func(p string) {
		pages.SwitchToPage(p)
	}

	pages.AddPage(
		"welcome",
		app.Welcome(nextSlide, a.StopApp),
		true,
		true,
	)

	pages.AddPage(
		"start",
		app.Start(nextSlide, a.StopApp, a.NewGame),
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
	if err := a.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}
