package rest

import (
	"github.com/gorilla/mux"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"hangman/services"
	"net/http/httptest"
	"time"
)

type MockHandlers interface {
	GetServer() *httptest.Server
	OnNewGameTimeout()
	OnNewGameReturn(id string)
	OnGuessReturn(id string, char rune, correct bool, missesLeft int, gameOver bool)
	AssertExpectations(t mock.TestingT)
	Close()
	URL() string
}

func NewMockHandlers() MockHandlers {
	gs := new(services.GameServiceMock)
	r := mux.NewRouter()
	a := NewGameServer(r, gs)
	a.InitialiseHandlers()
	return &httpServer{
		server: httptest.NewServer(r),
		gs:     gs,
	}
}

type httpServer struct {
	server *httptest.Server
	gs     *services.GameServiceMock
}

func (h *httpServer) GetServer() *httptest.Server {
	return h.server
}

func (h *httpServer) OnNewGameReturn(expectedId string) {
	h.gs.On("NewGame", mock.Anything).Return(expectedId)
}

func (h *httpServer) OnNewGameTimeout() {
	h.gs.On("NewGame", mock.Anything).After(1 * time.Second).Return("timeout")
}

func (h *httpServer) OnGuessReturn(id string, char rune, correct bool, missesLeft int, gameOver bool) {
	h.gs.On("Guess", id, char).Return(correct, missesLeft, gameOver)
}

func (h *httpServer) AssertExpectations(t mock.TestingT) {
	Expect(
		h.gs.AssertExpectations(t)).
		To(BeTrue())
}

func (h *httpServer) VerifyNewGameCalled(t mock.TestingT) {
	Expect(
		h.gs.AssertCalled(t, "Stop")).
		To(BeTrue())
}

func (h *httpServer) Close() {
	h.server.Close()
}

func (h *httpServer) URL() string {
	return h.server.URL
}
