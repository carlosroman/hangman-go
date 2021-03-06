package rest

import "net/http"

func New(cfg *Config) ApiClient {
	return &client{
		hc:      cfg.HTTPClient,
		baseURL: *cfg.BaseURL,
	}
}

type ApiClient interface {
	NewGame(name string, difficulty string) (gid string, err error)
	MakeGuess(gid string, guess rune) (correct bool, missesLeft int8, gameOver bool, currentWord []rune, err error)
}

type client struct {
	hc      *http.Client
	baseURL string
}
