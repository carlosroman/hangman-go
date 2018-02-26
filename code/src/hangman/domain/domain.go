package domain

import "sync"

type Status int

type State struct {
	sync.RWMutex
	Id     int
	Status Status
	Misses []rune
}

const (
	IN_PROGRESS Status = 1 + iota
	FINISHED
)
