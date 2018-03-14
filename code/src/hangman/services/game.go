package services

import (
	"hangman/domain"
	"hangman/services/wordstore"
	"sync"
)

type GameService interface {
	NewGame(d domain.Difficulty) int
	Guess(id int, char rune) (bool, int)
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

func (gs *inMemoryGameService) Guess(id int, char rune) (bool, int) {
	gs.RLock()
	defer gs.RUnlock()
	gs.games[id].Lock()
	defer gs.games[id].Unlock()

	gs.games[id].Guesses = append(gs.games[id].Guesses, char)
	f := gs.games[id].Word.Contains(char)
	if !f {
		gs.games[id].Misses += 1
	}
	return f, 8 - gs.games[id].Misses
}

func (gs *inMemoryGameService) GetGame(id int) domain.State {
	return gs.games[id]
}
