package rest

import (
	"errors"
	"hangman/domain"
	"strings"
)

type Guess struct {
	Guess rune `json:"guess"`
}

type GuessResponse struct {
	Correct    bool     `json:"correct"`
	MissesLeft int      `json:"missesLeft"`
	GameOver   bool     `json:"gameOver"`
	Letters    []string `json:"letters"`
}

type Difficulty string

const (
	VERY_EASY Difficulty = "VERY_EASY"
	EASY      Difficulty = "EASY"
	NORMAL    Difficulty = "NORMAL"
	HARD      Difficulty = "HARD"
	VERY_HARD Difficulty = "VERY_HARD"
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

func (d Difficulty) String() string {
	return string(d)
}

func GetDifficulty(d string) (Difficulty, error) {
	switch strings.TrimSpace(strings.ToUpper(d)) {
	case VERY_EASY.String():
		return VERY_EASY, nil
	case EASY.String():
		return EASY, nil
	case NORMAL.String():
		return NORMAL, nil
	case HARD.String():
		return HARD, nil
	case VERY_HARD.String():
		return VERY_HARD, nil
	default:
		return "", errors.New("invalid Difficulty")
	}
}

type NewGame struct {
	Difficulty Difficulty `json:"difficulty"`
}
