package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"hangman/services"
	"net/http"
	"strconv"
)

type GameServer interface {
	InitialiseHandlers()
}

func NewGameServer(r  *mux.Router, gs services.GameService) GameServer {
	return &App{r, gs}
}

type App struct {
	r  *mux.Router
	gs services.GameService
}

func (a *App) InitialiseHandlers() {
	a.r.HandleFunc("/game", a.handleNewGame).Methods("POST")
	a.r.HandleFunc("/game/{id:[0-9]+}/guess", a.handleGuess).Methods("POST")
}

func (a *App) handleNewGame(w http.ResponseWriter, r *http.Request) {
	id := a.gs.NewGame()
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("/game/%d", id))
}

func (a *App) handleGuess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var g Guess
	decoder.Decode(&g) // todo handle error
	a.gs.Guess(int(id), g.Guess)
	w.WriteHeader(http.StatusOK)
}
