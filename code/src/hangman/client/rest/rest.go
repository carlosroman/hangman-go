package rest

import "net/http"

func New(cfg *Config) ApiClient {
	return &client{
		c: cfg.HTTPClient,
	}
}

type ApiClient interface {
	NewGame(name string, difficulty string) (string, error)
}

type client struct {
	c *http.Client
}
