package app

import (
	"github.com/rivo/tview"
	"hangman/client/rest"
	"net/http"
	"sync/atomic"
)

type App interface {
	StopApp()
	NewGame(name string, difficulty string) error
	SetRoot(root tview.Primitive, fullscreen bool) App
	Run() error
	GetGameId() string
}

type app struct {
	ta     tviewApplication
	ac     rest.ApiClient
	gameId atomic.Value
}

func (a *app) StopApp() {
	a.ta.Stop()
}

func (a *app) NewGame(name string, difficulty string) error {
	id, err := a.ac.NewGame(name, difficulty)
	if err == nil {
		a.gameId.Store(id)
	}
	return err
}

func (a *app) SetRoot(root tview.Primitive, fullscreen bool) App {
	a.ta.SetRoot(root, fullscreen)
	return a
}

func (a *app) Run() error {
	return a.ta.Run()
}

func (a *app) GetGameId() string {
	return a.gameId.Load().(string)
}

func NewApp(baseUrl *string) App {
	cfg := rest.NewConfig()
	cfg.BaseURL = baseUrl
	cfg.HTTPClient = http.DefaultClient
	return &app{
		ta: tview.NewApplication(),
		ac: rest.New(cfg),
	}
}

func NewConfiguredApp(ta tviewApplication, ac rest.ApiClient) App {
	return &app{
		ta: ta,
		ac: ac,
	}
}

type tviewApplication interface {
	//SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *tview.Application
	//GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey
	Run() error
	Stop()
	//Suspend(f func()) bool
	//Draw() *tview.Application
	//SetBeforeDrawFunc(handler func(screen tcell.Screen) bool) *tview.Application
	//GetBeforeDrawFunc() func(screen tcell.Screen) bool
	//SetAfterDrawFunc(handler func(screen tcell.Screen)) *tview.Application
	//GetAfterDrawFunc() func(screen tcell.Screen)
	SetRoot(root tview.Primitive, fullscreen bool) *tview.Application
	//ResizeToFullScreen(p tview.Primitive) *tview.Application
	//SetFocus(p tview.Primitive) *tview.Application
	//GetFocus() tview.Primitive
}
