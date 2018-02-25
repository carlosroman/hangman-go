package services

import (
	"hangman/domain"
	"sync"
)

type GameService interface {
	newGame() int
}

type inMemoryGameService struct {
	games []domain.State
	mux   sync.RWMutex
}

func NewGameService() GameService {
	return &inMemoryGameService{}
}

func (gs *inMemoryGameService) newGame() int {
	gs.mux.Lock()
	defer gs.mux.Unlock()
	gs.games = append(gs.games, domain.State{})
	return len(gs.games) - 1
}
