package services

import (
	"fmt"
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
		want int
	}{
		{name: " New game results", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := new(wordstore.StoreMock)
			s.On("GetWord", mock.AnythingOfType("domain.Difficulty")).Return("word", nil).Once()
			if got := NewGameService(s).NewGame(domain.NORMAL); got != tt.want {
				t.Errorf("newGame() = %v, want %v", got, tt.want)
			}
			assertWorsdStoreCalledCorrectly(t, s, domain.NORMAL)
		})
	}
}

func Test_NewGames(t *testing.T) {
	ws := wordstore.NewMock()
	ws.On("GetWord", mock.AnythingOfType("domain.Difficulty")).Return("word", nil).Twice()
	gs := NewGameService(ws)
	_ = gs.NewGame(domain.EASY)
	got := gs.NewGame(domain.HARD)
	assert.Equal(t, 1, got, "Should only be second game")
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
		name   string
		guess  rune
		expect []rune
	}{
		{name: "First", guess: 'a', expect: []rune{'a'}},
		{name: "Second", guess: 'b', expect: []rune{'a', 'b'}},
		{name: "Third", guess: 'a', expect: []rune{'a', 'b', 'a'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.False(t, gs.Guess(id, tt.guess))
			assert.Equal(t, gs.GetGame(id).Misses, tt.expect)
		})
	}
}
