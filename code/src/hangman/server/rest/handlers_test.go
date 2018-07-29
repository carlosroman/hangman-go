package rest_test

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"hangman/server/rest"
	"io"
	"net/http"
)

var _ = Describe("Handlers", func() {

	var (
		mh rest.MockHandlers
	)
	Describe("Server endpoints", func() {

		BeforeEach(func() {
			mh = rest.NewMockHandlers()
		})

		AfterEach(func() {
			mh.Close()
		})

		Context("when handlers loaded", func() {

			It("should allow the creation of new game", func() {

				mh.OnNewGameReturn("123e4567-e89b-12d3-a456-426655440000")

				b := rest.NewGame{Difficulty: rest.NORMAL}
				bs, err := json.Marshal(b)
				Expect(err).To(Succeed())
				body := bytes.NewReader(bs)

				req, err := http.NewRequest("POST", mh.URL()+"/game", body)
				Expect(err).Should(Succeed())
				req.Header.Set("Content-Type", "application/json")

				resp, err := http.DefaultClient.Do(req)
				Expect(err).Should(Succeed())
				defer resp.Body.Close()
				mh.AssertExpectations(GinkgoT())

				Expect(resp.StatusCode).To(Equal(201))
				Expect(resp.Header.Get("Location")).To(Equal("/game/123e4567-e89b-12d3-a456-426655440000"))
			})

			It("should make a correct guess", func() {

				mh.OnGuessReturn("123e4567-e89b-12d3-a456-426655440000", 'a', true, 7331)

				b := rest.Guess{Guess: 'a'}
				bs, err := json.Marshal(b)
				Expect(err).To(Succeed())
				body := bytes.NewReader(bs)

				req, err := http.NewRequest("POST", mh.URL()+"/game/123e4567-e89b-12d3-a456-426655440000/guess", body)
				Expect(err).To(Succeed())
				req.Header.Set("Content-Type", "application/json")

				resp, err := http.DefaultClient.Do(req)
				Expect(err).Should(Succeed())
				defer resp.Body.Close()
				mh.AssertExpectations(GinkgoT())

				Expect(resp.StatusCode).To(Equal(200))
				actual, err := bytesToString(resp.Body)
				Expect(err).To(Succeed())
				expected, err := jsonToString(rest.GuessResponse{
					Correct:     true,
					GuessesLeft: 7331,
				})
				Expect(err).Should(Succeed())
				Expect(actual).Should(MatchJSON(expected))
			})

			It("should make a bad guess", func() {

				mh.OnGuessReturn("123e4567-e89b-12d3-a456-426655440000", 'a', false, 1337)

				b := rest.Guess{Guess: 'a'}
				bs, err := json.Marshal(b)
				Expect(err).To(Succeed())
				body := bytes.NewReader(bs)

				req, err := http.NewRequest("POST", mh.URL()+"/game/123e4567-e89b-12d3-a456-426655440000/guess", body)
				Expect(err).To(Succeed())
				req.Header.Set("Content-Type", "application/json")

				resp, err := http.DefaultClient.Do(req)
				Expect(err).To(Succeed())
				defer resp.Body.Close()
				mh.AssertExpectations(GinkgoT())

				Expect(resp.StatusCode).To(Equal(200))
				actual, err := bytesToString(resp.Body)
				Expect(err).To(Succeed())
				expected, err := jsonToString(rest.GuessResponse{
					Correct:     false,
					GuessesLeft: 1337,
				})
				Expect(err).To(Succeed())
				Expect(actual).Should(MatchJSON(expected))
			})
		})
	})

})

func jsonToString(obj interface{}) (string, error) {
	bs, err := json.Marshal(obj)
	return string(bs), err
}

func bytesToString(rc io.ReadCloser) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(rc)
	return buf.String(), err
}
