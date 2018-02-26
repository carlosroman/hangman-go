package services

import (
	"hangman/domain"
	"sync"
)

type GameService interface {
	NewGame() int
	Guess(id int, char rune) bool
	GetGame(id int) domain.State
}

type inMemoryGameService struct {
	sync.RWMutex
	games []domain.State
}

func NewGameService() GameService {
	return &inMemoryGameService{}
}

func (gs *inMemoryGameService) NewGame() int {
	gs.Lock()
	defer gs.Unlock()
	gs.games = append(gs.games, domain.State{Id: len(gs.games)})
	return len(gs.games) - 1
}

func (gs *inMemoryGameService) Guess(id int, char rune) bool {
	gs.RLock()
	defer gs.RUnlock()
	gs.games[id].Lock()
	defer gs.games[id].Unlock()
	gs.games[id].Misses = append(gs.games[id].Misses, char)
	return false
}

func (gs *inMemoryGameService) GetGame(id int) domain.State {
	return gs.games[id]
}
