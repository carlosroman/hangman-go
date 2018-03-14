package handlers

import "hangman/domain"

type Guess struct {
	Guess rune `json:"guess"`
}

type GuessResponse struct {
	Correct     bool `json:"correct"`
	GuessesLeft int  `json:"guessesLeft"`
	//Letters     []rune `json:"letters"`
}

type NewGame struct {
	Difficulty domain.Difficulty `json:"difficulty"`
}
