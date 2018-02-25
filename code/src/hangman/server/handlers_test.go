package main

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"hangman/services"
	"net/http"
	"net/http/httptest"
)

func TestHandlers(t *testing.T) {
	r := mux.NewRouter()
	gs := new(services.GameServiceMock)
	a := App{
		gs: gs,
		r:  r,
	}

	a.initialiseHandlers()

	tests := []struct {
		name       string
		method     string
		statusCode int
		path       string
	}{
		{
			name:       "Create new game",
			method:     "POST",
			statusCode: 201,
			path:       "/game",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs.On("NewGame").Return(2)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			r.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}
