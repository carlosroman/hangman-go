package services

import "hangman/domain"

var games []domain.State

func newGame() int {
	games = append(games, domain.State{})
	return len(games) - 1
}