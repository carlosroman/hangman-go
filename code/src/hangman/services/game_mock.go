package services

import (
	"github.com/stretchr/testify/mock"
	"hangman/domain"
)

type GameServiceMock struct {
	mock.Mock
}

func (gs *GameServiceMock) NewGame(d domain.Difficulty) string {
	args := gs.Called(d)
	return args.String(0)
}

func (gs *GameServiceMock) Guess(id string, char rune) (correct bool, missesLeft int, gameOver bool) {
	args := gs.Called(id, char)
	return args.Bool(0), args.Int(1), args.Bool(2)
}

func (gs *GameServiceMock) GetGame(id string) domain.State {
	args := gs.Called(id)
	return args.Get(0).(domain.State)
}
