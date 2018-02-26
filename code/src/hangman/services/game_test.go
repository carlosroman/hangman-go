package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_newGame(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: " New game results", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGameService().NewGame(); got != tt.want {
				t.Errorf("newGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewGames(t *testing.T) {
	gs := NewGameService()
	_ = gs.NewGame()
	got := gs.NewGame()
	assert.Equal(t, 1, got, "Should only be second game")
}
func Test_GetGame(t *testing.T) {
	gs := NewGameService()
	tests := []struct {
		name string
	}{
		{name: "One"},
		{name: "Two"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := gs.NewGame()
			g := gs.GetGame(id)
			assert.Equal(t, id, g.Id)
		})
	}
}

func Test_Guess(t *testing.T) {
	gs := NewGameService()
	id := gs.NewGame()
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
