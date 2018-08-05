package domain

import (
	"sync"
)

type Status int

type State struct {
	*sync.RWMutex
	Id      string
	Status  Status
	Misses  int
	Guesses []rune
	Word    Word
}

const (
	IN_PROGRESS Status = iota + 1
	FINISHED
)

type Difficulty int

const (
	VERY_EASY Difficulty = iota + 1
	EASY
	NORMAL
	HARD
	VERY_HARD
)

func (d Difficulty) String() string {
	names := [...]string{
		"very easy",
		"easy",
		"normal",
		"hard",
		"very hard",
	}
	if d < VERY_EASY || d > VERY_HARD {
		return "unknown"
	}
	return names[d-1]
}

func NewWord(word string, difficulty Difficulty) Word {
	w := Word{
		Letters:    []rune(word),
		Difficulty: difficulty,
	}
	lg := make([]rune, len(w.Letters))
	for i := range lg {
		lg[i] = '_'
	}
	w.LetterGuessed = lg
	return w
}

type Word struct {
	Letters       []rune
	Difficulty    Difficulty
	LetterGuessed []rune
}

func (w *Word) Contains(l rune) (correct bool) {
	for i, value := range w.Letters {
		if value == l {
			w.LetterGuessed[i] = l
			correct = true
		}
	}
	return correct
}
