package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	srest "hangman/server/rest"
	"net/http"
	"strings"
	"unicode/utf8"
)

func newGameRequest(baseURL string, d srest.Difficulty) (req *http.Request, err error) {
	p := srest.NewGame{Difficulty: d}
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(p); err != nil {
		return req, err
	}
	req, err = http.NewRequest("POST", baseURL+"/game", b)
	if err != nil {
		return req, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	return req, err
}

func parseNewGameResponse(resp *http.Response) (gid string, err error) {
	if resp.StatusCode != 201 {
		return gid, errors.New("new game not created")
	}

	l := resp.Header.Get("Location")
	gid = strings.Split(l, "/")[2]

	return gid, err
}

func newGuessRequest(baseURL string, gid string, guess rune) (req *http.Request, err error) {
	g := srest.Guess{Guess: guess}
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(g); err != nil {
		return req, err
	}
	req, err = http.NewRequest("POST", fmt.Sprintf("%s/game/%s/guess", baseURL, gid), b)
	if err != nil {
		return req, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	return req, err
}

func parseNewGuessResponse(resp *http.Response) (correct bool, missesLeft int8, gameOver bool, currentWord []rune, err error) {
	if resp.StatusCode != 200 {
		fmt.Println(fmt.Sprintf("Status code was '%d'", resp.StatusCode))
		return correct, missesLeft, gameOver, currentWord, errors.New("guess could not be made")
	}

	var gr srest.GuessResponse

	if err := json.NewDecoder(resp.Body).Decode(&gr); err != nil {
		return correct, missesLeft, gameOver, currentWord, err
	}
	currentWord = make([]rune, len(gr.Letters))
	for i, l := range gr.Letters {
		currentWord[i], _ = utf8.DecodeRuneInString(l)
	}
	return gr.Correct, int8(gr.MissesLeft), gr.GameOver, currentWord, err
}
