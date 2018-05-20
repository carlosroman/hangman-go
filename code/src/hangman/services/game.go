package services

import (
	"github.com/satori/go.uuid"
	"hangman/domain"
	"hangman/services/wordstore"
	"sync"
)

type GameService interface {
	NewGame(d domain.Difficulty) string
	Guess(id string, char rune) (bool, int)
	GetGame(id string) domain.State
}

type inMemoryGameService struct {
	sync.RWMutex
	games map[string]*domain.State
	w     wordstore.Store
}

func NewGameService(words wordstore.Store) GameService {
	return &inMemoryGameService{w: words, games: make(map[string]*domain.State)}
}

func (gs *inMemoryGameService) NewGame(d domain.Difficulty) string {
	gs.Lock()
	defer gs.Unlock()
	w, _ := gs.w.GetWord(d)
	wd := domain.Word{
		Letters:    []rune(w),
		Difficulty: d,
	}
	id := uuid.NewV4().String()
	gs.games[id] = &domain.State{Id: id, Word: wd}
	return id
}

func (gs *inMemoryGameService) Guess(id string, char rune) (bool, int) {
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

func (gs *inMemoryGameService) GetGame(id string) domain.State {
	return *gs.games[id]
}
