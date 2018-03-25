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

type Difficulty string

const (
	VERY_EASY Difficulty = "VERY_EASY"
	EASY                 = "EASY"
	NORMAL               = "NORMAL"
	HARD                 = "HARD"
	VERY_HARD            = "VERY_HARD"
)

func (d Difficulty) toDomainDifficulty() domain.Difficulty {
	switch d {
	case VERY_EASY:
		return domain.VERY_EASY
	case EASY:
		return domain.EASY
	case NORMAL:
		return domain.NORMAL
	case HARD:
		return domain.HARD
	case VERY_HARD:
		return domain.VERY_HARD
	default:
		return -1
	}
}

type NewGame struct {
	Difficulty Difficulty `json:"difficulty"`
}
