package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strings"
)

const logo = `
 __   __  _______  __    _  _______  __   __  _______  __    _ 
|  | |  ||   _   ||  |  | ||       ||  |_|  ||   _   ||  |  | |
|  |_|  ||  |_|  ||   |_| ||    ___||       ||  |_|  ||   |_| |
|       ||       ||       ||   | __ |       ||       ||       |
|       ||       ||  _    ||   ||  ||       ||       ||  _    |
|   _   ||   _   || | |   ||   |_| || ||_|| ||   _   || | |   |
|__| |__||__| |__||_|  |__||_______||_|   |_||__| |__||_|  |__|

`
const (
	subtitle = `Hangman - the classic word guessing game`
)

func Welcome(nextSlide func(p string), stopApp func()) tview.Primitive {

	// What's the size of the logo?
	lines := strings.Split(logo, "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := tview.NewTextView().
		SetTextColor(tcell.ColorGreen)

	fmt.Fprint(logoBox, logo)

	// Create a frameSubtitle for the subtitle and game prompt
	frameSt := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite)
	f := tview.NewForm().
		AddButton("Play", func() {
			nextSlide("start")
		}).
		AddButton("Quit", stopApp)

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, false).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, false).
		AddItem(frameSt, 0, 10, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 3, false).
			AddItem(f, 0, 1, true).
			AddItem(tview.NewBox(), 0, 3, false), logoHeight, 2, true)

	return flex
}
