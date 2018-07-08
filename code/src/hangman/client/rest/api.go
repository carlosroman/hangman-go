package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"hangman/server/handlers"
	"net/http"
	"strings"
)

func (c *client) NewGame(name string, difficulty string) (string, error) {

	d, err := handlers.GetDifficulty(difficulty)
	if err != nil {
		return "", err
	}
	p := handlers.NewGame{Difficulty: d}
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(p); err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.baseURL+"/game", b)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := c.hc.Do(req)
	defer func(r *http.Response) {
		if r != nil {
			r.Body.Close()
		}
	}(resp)

	var gid string
	if resp != nil {
		if resp.StatusCode != 201 {
			err = errors.New("new game not created")
		} else {
			l := resp.Header.Get("Location")
			gid = strings.Split(l, "/")[2]
		}
	}

	return gid, err
}
