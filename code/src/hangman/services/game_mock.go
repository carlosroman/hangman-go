package services

import (
	"github.com/stretchr/testify/mock"
)

type GameServiceMock struct {
	mock.Mock
}

func (gs *GameServiceMock) NewGame() int {
	args := gs.Called()
	return args.Int(0)
}
