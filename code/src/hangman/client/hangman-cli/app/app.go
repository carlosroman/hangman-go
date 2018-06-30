package app

import "github.com/rivo/tview"

type App interface {
	StopApp()
	NewGame(name string, difficulty string) error
	SetRoot(root tview.Primitive, fullscreen bool) App
	Run() error
}

type app struct {
	ta *tview.Application
}

func (a *app) StopApp() {
	a.ta.Stop()
}

func (a *app) NewGame(name string, difficulty string) error {
	return nil
}

func (a *app) SetRoot(root tview.Primitive, fullscreen bool) App {
	a.ta.SetRoot(root, fullscreen)
	return a
}

func (a *app) Run() error {
	return a.ta.Run()
}

func NewApp() App {
	return &app{tview.NewApplication()}
}
