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
	Guess(id string, char rune) (correct bool, missesLeft int, gameOver bool, letters []rune)
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
	wd := domain.NewWord(w, d)
	id := uuid.NewV4().String()

	gs.games[id] = &domain.State{Id: id, Word: wd, RWMutex: &sync.RWMutex{}}
	return id
}

func (gs *inMemoryGameService) Guess(id string, char rune) (bool, int, bool, []rune) {
	gs.RLock()
	defer gs.RUnlock()
	gs.games[id].Lock()
	defer gs.games[id].Unlock()

	gs.games[id].Guesses = append(gs.games[id].Guesses, char)
	if gs.games[id].Misses == maxMisses {
		return false, maxMisses - gs.games[id].Misses, true, gs.games[id].Word.LetterGuessed
	}

	f := gs.games[id].Word.Contains(char)
	if !f {
		gs.games[id].Misses += 1
		if gs.games[id].Misses == maxMisses {
			return f, maxMisses - gs.games[id].Misses, true, gs.games[id].Word.LetterGuessed
		}
	}
	return f, maxMisses - gs.games[id].Misses, false, gs.games[id].Word.LetterGuessed
}

func (gs *inMemoryGameService) GetGame(id string) domain.State {
	return *gs.games[id]
}
