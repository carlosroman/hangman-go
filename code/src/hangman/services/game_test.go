package services

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hangman/domain"
	"hangman/services/wordstore"
	"testing"
)

func assertWorsdStoreCalledCorrectly(t *testing.T, s *wordstore.StoreMock, d domain.Difficulty) {
	assert.True(t,
		s.AssertCalled(t, "GetWord", domain.NORMAL),
		fmt.Sprintf("Expected word store to be called with Difficulty '%s'", domain.NORMAL))
}

func Test_newGame(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: " New game results"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := new(wordstore.StoreMock)
			s.On("GetWord", mock.AnythingOfType("domain.Difficulty")).Return("word", nil).Once()
			got := NewGameService(s).NewGame(domain.NORMAL)
			u, err := uuid.FromString(got)
			assert.NoError(t, err)
			assert.Equal(t, uuid.V4, u.Version())
			assertWorsdStoreCalledCorrectly(t, s, domain.NORMAL)
		})
	}
}

func Test_NewGames(t *testing.T) {
	ws := wordstore.NewMock()
	ws.On("GetWord", mock.AnythingOfType("domain.Difficulty")).Return("word", nil).Twice()
	gs := NewGameService(ws)
	one := gs.NewGame(domain.EASY)
	got := gs.NewGame(domain.HARD)
	assert.NotEqual(t, one, got, "Should both be different")
}

func Test_GetGame(t *testing.T) {
	ws := wordstore.NewMock()
	ws.On("GetWord", mock.AnythingOfType("domain.Difficulty")).Return("word", nil).Twice()
	gs := NewGameService(ws)
	tests := []struct {
		name       string
		difficulty domain.Difficulty
	}{
		{name: "One", difficulty: domain.NORMAL},
		{name: "Two", difficulty: domain.VERY_HARD},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := gs.NewGame(tt.difficulty)
			g := gs.GetGame(id)
			assert.Equal(t, id, g.Id)
			assertWorsdStoreCalledCorrectly(t, ws, tt.difficulty)
		})
	}
}

func Test_Guess(t *testing.T) {
	ws := wordstore.NewMock()
	ws.On("GetWord", mock.AnythingOfType("domain.Difficulty")).Return("word", nil).Once()
	gs := NewGameService(ws)
	id := gs.NewGame(domain.VERY_HARD)
	assert.Empty(t, gs.GetGame(id).Misses)

	tests := []struct {
		name              string
		guess             rune
		expectGuesses     []rune
		expectFound       bool
		expectedGuessLeft int
	}{
		{name: "First", guess: 'w', expectGuesses: []rune{'w'}, expectFound: true, expectedGuessLeft: 8},
		{name: "Second", guess: 'b', expectGuesses: []rune{'w', 'b'}, expectFound: false, expectedGuessLeft: 7},
		{name: "Third", guess: 'd', expectGuesses: []rune{'w', 'b', 'd'}, expectFound: true, expectedGuessLeft: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, m := gs.Guess(id, tt.guess)
			assert.Equal(t, r, tt.expectFound, fmt.Sprintf("Expected Found to be %s", tt.expectFound))
			assert.Equal(t, m, tt.expectedGuessLeft, fmt.Sprintf("Expected Guess Left to be %s", tt.expectedGuessLeft))
			assert.Equal(t, gs.GetGame(id).Guesses, tt.expectGuesses)
		})
	}
}
