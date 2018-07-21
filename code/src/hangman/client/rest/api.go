package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	srest "hangman/server/rest"
	"net/http"
	"strings"
)

func (c *client) NewGame(name string, difficulty string) (string, error) {
	var gid string

	d, err := srest.GetDifficulty(difficulty)
	if err != nil {
		return gid, err
	}
	p := srest.NewGame{Difficulty: d}
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(p); err != nil {
		return gid, err
	}

	req, err := http.NewRequest("POST", c.baseURL+"/game", b)
	if err != nil {
		return gid, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := c.hc.Do(req)
	defer func(r *http.Response) {
		if r != nil {
			r.Body.Close()
		}
	}(resp)

	if err != nil {
		return gid, err
	}

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
