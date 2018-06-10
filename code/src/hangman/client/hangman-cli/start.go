package main

import (
	"github.com/rivo/tview"
)

func Start(nextSlide func(p string), stopApp func()) tview.Primitive {
	d := []string{
		"Very easy",
		"Easy",
		"Normal",
		"Hard",
		"Very hard",
	}
	f := tview.NewForm().
		AddInputField("Nickname:", "", 20, nil, nil).
		AddDropDown("Difficulty:", d, 0, nil).
		AddButton("Start", func() {
			// todo: validate form
			// todo: call create new game
			nextSlide("go")
		}).
		AddButton("Quit", stopApp)
	f.SetBorder(true).SetTitle("New Game")
	return f
}
