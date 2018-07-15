package app

import (
	"github.com/rivo/tview"
	"fmt"
)

func Start(nextSlide func(p string),
	stopApp func(),
	newGame func(name string, difficulty string) error) tview.Primitive {

	d := []string{
		"Very easy",
		"Easy",
		"Normal",
		"Hard",
		"Very hard",
	}

	f := tview.NewForm().
		AddInputField("Nickname:", "", 20, nil, nil).
		AddDropDown("Difficulty:", d, 0, nil)

	f.AddButton("Start", func() {
		nic := f.GetFormItemByLabel("Nickname:").(*tview.InputField).GetText()
		_, sd := f.GetFormItemByLabel("Difficulty:").(*tview.DropDown).GetCurrentOption()
		if err := newGame(
			nic,
			sd); err != nil {
				fmt.Println(err)
		}
		nextSlide("go")
	}).AddButton("Quit", stopApp)
	f.SetBorder(true).SetTitle("New Game")
	return f
}
