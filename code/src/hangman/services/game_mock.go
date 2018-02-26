package services

import (
	"github.com/stretchr/testify/mock"
	"hangman/domain"
)

type GameServiceMock struct {
	mock.Mock
}

func (gs *GameServiceMock) NewGame() int {
	args := gs.Called()
	return args.Int(0)
}

func (gs *GameServiceMock) Guess(id int, char rune) bool {
	args := gs.Called(id, char)
	return args.Bool(0)
}

func (gs *GameServiceMock) GetGame(id int) domain.State {
	args := gs.Called(id)
	return args.Get(0).(domain.State)
}
