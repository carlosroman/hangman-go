package rest_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	cli "hangman/client/rest"
	"hangman/server/rest"
	"math/rand"
	"net/http"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var _ = Describe("Api", func() {

	h := &http.Client{
		Timeout: 1 * time.Second,
	}

	var (
		mh rest.MockHandlers
		c  cli.ApiClient
	)

	Describe("Bad config", func() {
		It("should fail to create a new game", func() {
			cfg := cli.NewConfig()
			cfg.HTTPClient = h
			burl := "http://us\ner:pass\nword@foo.com"
			cfg.BaseURL = &burl
			c = cli.New(cfg)
			_, err := c.NewGame("Bob", "easy")
			fmt.Fprintln(GinkgoWriter, err)
			Expect(err).ShouldNot(Succeed())
		})
		It("should fail to make a guess", func() {
			cfg := cli.NewConfig()
			cfg.HTTPClient = h
			burl := "http://us\ner:pass\nword@foo.com"
			cfg.BaseURL = &burl
			c = cli.New(cfg)
			_, _, _, err := c.MakeGuess(randomString(12), 'a')
			fmt.Fprintln(GinkgoWriter, err)
			Expect(err).ShouldNot(Succeed())
		})

		It("should error if http client returns a 404", func() {
			mh = rest.NewMockHandlers()
			defer mh.Close()
			cfg := cli.NewConfig()
			cfg.HTTPClient = h
			burl := mh.URL() + "/404"
			cfg.BaseURL = &burl
			c = cli.New(cfg)
			_, err := c.NewGame("Bob", "easy")
			fmt.Fprintln(GinkgoWriter, err)
			Expect(err).ShouldNot(Succeed())
		})
	})

	Describe("New game", func() {

		BeforeEach(func() {
			mh = rest.NewMockHandlers()
			cfg := cli.NewConfig()
			cfg.HTTPClient = h
			burl := mh.URL()
			cfg.BaseURL = &burl
			c = cli.New(cfg)
		})

		AfterEach(func() {
			mh.Close()
		})

		It("should post correct data", func() {
			mh.OnNewGameReturn("game-id")
			id, err := c.NewGame("Bob", "easy")
			mh.AssertExpectations(GinkgoT())
			Expect(err).Should(Succeed())
			Expect(id).To(Equal("game-id"))
		})

		It("should error for wrong difficulty", func() {
			_, err := c.NewGame("Bob", "incorrect-difficulty")
			fmt.Fprintln(GinkgoWriter, err)
			Expect(err).ShouldNot(Succeed())
		})

		It("should error for http client issues like timeout", func() {
			mh.OnNewGameTimeout()
			_, err := c.NewGame("Bob", "easy")
			fmt.Fprintln(GinkgoWriter, err)
			Expect(err).ShouldNot(Succeed())
		})
	})

	Describe("Get game status", func() {

		BeforeEach(func() {
			mh = rest.NewMockHandlers()
			cfg := cli.NewConfig()
			cfg.HTTPClient = h
			burl := mh.URL()
			cfg.BaseURL = &burl
			c = cli.New(cfg)
		})

		AfterEach(func() {
			mh.Close()
		})

		It("should return correct number of guesses left", func() {
			gid := uuid.NewV4().String()
			mh.OnGuessReturn(gid, 'a', true, 8, false)
			_, missesLeft, _, err := c.MakeGuess(gid, 'a')
			Expect(err).To(Succeed())
			Expect(missesLeft).To(Equal(int8(8)))
			mh.AssertExpectations(GinkgoT())
		})

		It("should return true when guess correct", func() {
			gid := uuid.NewV4().String()
			mh.OnGuessReturn(gid, 'a', true, 8, false)
			g, _, _, err := c.MakeGuess(gid, 'a')
			Expect(err).To(Succeed())
			Expect(g).To(BeTrue(), "Expect guess to be correct (true)")
			mh.AssertExpectations(GinkgoT())
		})

		It("should return false when guess incorrect", func() {
			gid := uuid.NewV4().String()
			mh.OnGuessReturn(gid, 'a', false, 8, false)
			g, _, _, err := c.MakeGuess(gid, 'a')
			Expect(err).To(Succeed())
			Expect(g).To(BeFalse(), "Expect guess to be incorrect (false)")
			mh.AssertExpectations(GinkgoT())
		})

		It("should return true when game over", func() {
			gid := uuid.NewV4().String()
			mh.OnGuessReturn(gid, 'a', false, 0, true)
			_, _, g, err := c.MakeGuess(gid, 'a')
			Expect(err).To(Succeed())
			Expect(g).To(BeTrue(), "Expect game over to be true")
			mh.AssertExpectations(GinkgoT())
		})

		It("should return false when game not over", func() {
			gid := uuid.NewV4().String()
			mh.OnGuessReturn(gid, 'a', false, 0, false)
			_, _, g, err := c.MakeGuess(gid, 'a')
			Expect(err).To(Succeed())
			Expect(g).To(BeFalse(), "Expect game over to be false")
			mh.AssertExpectations(GinkgoT())
		})

		It("should return error when guess could not be made", func() {
			gid := randomString(12)
			mh.OnGuessReturn(gid, 'a', true, 8, false)
			_, _, _, err := c.MakeGuess(gid, 'a')
			Expect(err).NotTo(Succeed())
		})
	})
})
