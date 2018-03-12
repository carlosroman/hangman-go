package services

import (
	"hangman/domain"
	"hangman/services/wordstore"
	"sync"
)

type GameService interface {
	NewGame(d domain.Difficulty) int
	Guess(id int, char rune) bool
	GetGame(id int) domain.State
}

type inMemoryGameService struct {
	sync.RWMutex
	games []domain.State
	w     wordstore.Store
}

func NewGameService(words wordstore.Store) GameService {
	return &inMemoryGameService{w: words}
}

func (gs *inMemoryGameService) NewGame(d domain.Difficulty) int {
	gs.Lock()
	defer gs.Unlock()
	w, _ := gs.w.GetWord(d)
	wd := domain.Word{
		Letters:    []rune(w),
		Difficulty: d,
	}
	gs.games = append(gs.games, domain.State{Id: len(gs.games), Word: wd})
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
