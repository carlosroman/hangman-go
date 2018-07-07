package rest

import (
	"net/http"
)

type Config struct {
	BaseURL    *string
	HTTPClient *http.Client
}

func NewConfig() *Config {
	return &Config{}
}
