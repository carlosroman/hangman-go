package rest_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	cli "hangman/client/rest"
	"hangman/server/rest"
	"net/http"
	"time"
)

var _ = Describe("Api", func() {

	h := &http.Client{
		Timeout: 1 * time.Second,
	}

	var (
		mh rest.MockHandlers
		c  cli.ApiClient
	)

	Describe("Bad config", func() {
		It("should post correct data", func() {
			cfg := cli.NewConfig()
			cfg.HTTPClient = h
			burl := "http://us\ner:pass\nword@foo.com"
			cfg.BaseURL = &burl
			c = cli.New(cfg)
			_, err := c.NewGame("Bob", "easy")
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
})
