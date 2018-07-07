package rest

import (
	"net/http"
)

type Config struct {
	HostName   *string
	HTTPClient *http.Client
}
