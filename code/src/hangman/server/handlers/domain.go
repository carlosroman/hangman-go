package handlers

import "hangman/domain"

type Guess struct {
	Guess rune `json:"guess"`
}

type NewGame struct {
	Difficulty domain.Difficulty `json:"difficulty"`
}
