package rest

import "net/http"

func New(cfg *Config) ApiClient {
	return &client{
		hc:      cfg.HTTPClient,
		baseURL: *cfg.BaseURL,
	}
}

type ApiClient interface {
	NewGame(name string, difficulty string) (string, error)
}

type client struct {
	hc      *http.Client
	baseURL string
}
