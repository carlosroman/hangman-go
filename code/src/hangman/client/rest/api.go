package rest

import (
	srest "hangman/server/rest"
	"net/http"
)

func (c *client) NewGame(name string, difficulty string) (string, error) {
	var gid string

	d, err := srest.GetDifficulty(difficulty)
	if err != nil {
		return gid, err
	}

	req, err := newGameRequest(c.baseURL, d)
	if err != nil {
		return gid, err
	}

	resp, err := c.hc.Do(req)
	defer func(r *http.Response) {
		if r != nil {
			r.Body.Close()
		}
	}(resp)

	if err != nil {
		return gid, err
	}

	return parseNewGameResponse(resp)
}

func (c *client) MakeGuess(gid string, guess rune) (correct bool, missesLeft int8, gameOver bool, currentWord []rune, err error) {
	req, err := newGuessRequest(c.baseURL, gid, guess)
	if err != nil {
		return correct, missesLeft, gameOver, currentWord, err
	}

	resp, err := c.hc.Do(req)
	defer func(r *http.Response) {
		if r != nil {
			r.Body.Close()
		}
	}(resp)

	if err != nil {
		return correct, missesLeft, gameOver, currentWord, err
	}

	return parseNewGuessResponse(resp)
}
