package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	pages := tview.NewPages()

	nextSlide := func(p string) {
		pages.SwitchToPage(p)
	}

	stopApp := func() {
		app.Stop()
	}

	newGame := func(name string, difficulty string) error {
		return nil
	}

	pages.AddPage(
		"welcome",
		Welcome(nextSlide, stopApp),
		true,
		true,
	)

	pages.AddPage(
		"start",
		Start(nextSlide, stopApp, newGame),
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
	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}
