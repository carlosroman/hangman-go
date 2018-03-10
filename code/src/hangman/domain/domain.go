package domain

import (
	"sync"
)

type Status int

type State struct {
	sync.RWMutex
	Id     int
	Status Status
	Misses []rune
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

type Word struct {
	Letters    []rune
	Difficulty Difficulty
}
