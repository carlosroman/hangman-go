package handlers

import (
	"testing"

	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hangman/domain"
	"hangman/services"
	"io"
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

	a.InitialiseHandlers()

	tests := []struct {
		name       string
		method     string
		statusCode int
		path       string
		body       interface{}
	}{
		{
			name:       "Create new game",
			method:     "POST",
			statusCode: 201,
			path:       "/game",
			body:       NewGame{Difficulty: domain.NORMAL},
		},
		{
			name:       "Make a guess",
			method:     "POST",
			statusCode: 200,
			path:       "/game/2/guess",
			body:       Guess{Guess: 'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs.On("NewGame", mock.Anything).Return(2)
			gs.On("Guess", 2, 'a').Return(true)

			rec := httptest.NewRecorder()

			req := funcName(tt)
			r.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}
func funcName(tt struct {
	name       string
	method     string
	statusCode int
	path       string
	body       interface{}
}) *http.Request {
	var pl io.Reader
	if tt.body != nil {
		b, _ := json.Marshal(tt.body)
		pl = bytes.NewReader(b)
	}
	req, _ := http.NewRequest(tt.method, tt.path, pl)
	return req
}
