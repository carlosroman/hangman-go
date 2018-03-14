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
	"hangman/utils"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var log = utils.Logger()

func TestHandlers(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		statusCode int
		path       string
		body       interface{}
		resp       string
		setup      func(gs *services.GameServiceMock)
	}{
		{
			name:       "Create new game",
			method:     "POST",
			statusCode: 201,
			path:       "/game",
			body:       NewGame{Difficulty: domain.NORMAL},
			setup: func(gs *services.GameServiceMock) {
				gs.On("NewGame", mock.Anything).Return(2)
			},
		},
		{
			name:       "Make a successful guess",
			method:     "POST",
			statusCode: 200,
			path:       "/game/2/guess",
			body:       Guess{Guess: 'a'},
			resp:       "{\"correct\":true,\"guessesLeft\":7331}\n",
			setup: func(gs *services.GameServiceMock) {
				gs.On("NewGame", mock.Anything).Return(2)
				gs.On("Guess", 2, 'a').Return(true, 7331)
			},
		},
		{
			name:       "Make a bad guess",
			method:     "POST",
			statusCode: 200,
			path:       "/game/2/guess",
			body:       Guess{Guess: 'a'},
			resp:       "{\"correct\":false,\"guessesLeft\":1337}\n",
			setup: func(gs *services.GameServiceMock) {
				gs.On("NewGame", mock.Anything).Return(2)
				gs.On("Guess", 2, 'a').Return(false, 1337)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := mux.NewRouter()

			gs := new(services.GameServiceMock)
			a := App{
				gs: gs,
				r:  r,
			}
			a.InitialiseHandlers()

			tt.setup(gs)

			rec := httptest.NewRecorder()

			req := funcName(tt)
			r.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
			if tt.resp != "" {
				b, _ := ioutil.ReadAll(rec.Body)
				assert.Equal(t, tt.resp, string(b))
			}
		})
	}
}

func funcName(tt struct {
	name       string
	method     string
	statusCode int
	path       string
	body       interface{}
	resp       string
	setup      func(gs *services.GameServiceMock)
}) *http.Request {
	var pl io.Reader
	if tt.body != nil {
		b, _ := json.Marshal(tt.body)
		log.Info(string(b[:]))
		pl = bytes.NewReader(b)
	}
	req, _ := http.NewRequest(tt.method, tt.path, pl)
	return req
}
