package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
			if got := newGame(); got != tt.want {
				t.Errorf("newGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newGames(t *testing.T) {
	_ = newGame()
	got := newGame()
	assert.Equal(t, 1, got, "Should only be second game")
}