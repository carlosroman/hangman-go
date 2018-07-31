package services

import (
	"github.com/satori/go.uuid"
	"hangman/domain"
	"hangman/services/wordstore"
	"sync"
)

const (
	maxMisses = 8
)

type GameService interface {
	NewGame(d domain.Difficulty) (gameId string)
	Guess(id string, char rune) (correct bool, missesLeft int, gameOver bool)
	GetGame(id string) (game domain.State)
}

type inMemoryGameService struct {
	sync.RWMutex
	games map[string]*domain.State
	w     wordstore.Store
}

func NewGameService(words wordstore.Store) GameService {
	return &inMemoryGameService{
		w:     words,
		games: make(map[string]*domain.State),
	}
}

func (gs *inMemoryGameService) NewGame(d domain.Difficulty) string {
	gs.Lock()
	defer gs.Unlock()
	w, _ := gs.w.GetWord(d) // todo: error handle
	wd := domain.Word{
		Letters:    []rune(w),
		Difficulty: d,
	}
	id := uuid.NewV4().String()

	gs.games[id] = &domain.State{Id: id, Word: wd, RWMutex: &sync.RWMutex{}}
	return id
}

func (gs *inMemoryGameService) Guess(id string, char rune) (bool, int, bool) {
	gs.RLock()
	defer gs.RUnlock()
	gs.games[id].Lock()
	defer gs.games[id].Unlock()

	gs.games[id].Guesses = append(gs.games[id].Guesses, char)
	if gs.games[id].Misses == maxMisses {
		return false, maxMisses - gs.games[id].Misses, true
	}

	f := gs.games[id].Word.Contains(char)
	if !f {
		gs.games[id].Misses += 1
		if gs.games[id].Misses == maxMisses {
			return f, maxMisses - gs.games[id].Misses, true
		}
	}
	return f, maxMisses - gs.games[id].Misses, false
}

func (gs *inMemoryGameService) GetGame(id string) domain.State {
	return *gs.games[id]
}
