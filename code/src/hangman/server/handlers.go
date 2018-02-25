package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"hangman/services"
	"net/http"
)

type App struct {
	r  *mux.Router
	gs services.GameService
}

func (a *App) initialiseHandlers() {
	a.r.HandleFunc("/game", a.handleNewGame).Methods("POST")
}

func (a *App) handleNewGame(w http.ResponseWriter, r *http.Request) {
	id := a.gs.NewGame()
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("/game/%d", id))
}
