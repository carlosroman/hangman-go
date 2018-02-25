package domain

type Status int

type State struct {
	id int
	status Status
}

const (
	IN_PROGRESS  Status = 1 + iota
	FINISHED
)