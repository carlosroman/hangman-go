package wordstore

import (
	"github.com/stretchr/testify/mock"
	"hangman/domain"
)

type StoreMock struct {
	mock.Mock
}

func (s *StoreMock) GetWord(d domain.Difficulty) (string, error) {
	args := s.Called(d)
	return args.String(0), args.Error(1)
}

func NewMock() *StoreMock {
	return new(StoreMock)
}
