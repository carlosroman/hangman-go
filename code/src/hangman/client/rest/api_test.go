package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"hangman/server/handlers"
	"io"
	"io/ioutil"
	"net/http"
)

var _ = Describe("Api", func() {

	var (
		rt *mockRoundTripper
		c  ApiClient
	)

	BeforeEach(func() {
		rt = &mockRoundTripper{}
		burl := "http://some.base/url"
		c = newTestApiClientrt(rt, burl)
	})

	Describe("New game", func() {

		It("should post correct data", func() {
			rt.givenResponseBackIs(201, "test", nil)
			_, err := c.NewGame("Bob", "easy")
			rt.AssertExpectations(GinkgoT())
			Expect(err).Should(Succeed())
			rt.assertRequest(
				"http://some.base/url/game",
				"POST",
				handlers.NewGame{Difficulty: handlers.EASY})
		})

		It("should save return game ID", func() {
			rt.givenResponseBackIs(201, "test", nil)
			gid, err := c.NewGame("Bob", "easy")
			rt.AssertExpectations(GinkgoT())
			Expect(err).Should(Succeed())
			Expect(gid).To(Equal("123e4567-e89b-12d3-a456-426655440000"))
		})

		It("should error for wrong difficulty", func() {
			_, err := c.NewGame("Bob", "incorrect-difficulty")
			rt.AssertExpectations(GinkgoT())
			Expect(err).ShouldNot(Succeed())
		})

		It("should error for creating a bad New Request", func() {
			badUrl := "http://us\ner:pass\nword@foo.com"
			c = newTestApiClientrt(rt, badUrl)
			_, err := c.NewGame("Bob", "easy")
			rt.AssertExpectations(GinkgoT())
			Expect(err).ShouldNot(Succeed())
		})

		It("should error if http client returns an error", func() {
			err := errors.New("some http error")
			rt.givenResponseBackIs(404, "test", err)
			_, err = c.NewGame("Bob", "easy")
			rt.AssertExpectations(GinkgoT())
			Expect(err).ShouldNot(Succeed())
		})
	})
})

func newTestApiClientrt(rt *mockRoundTripper, burl string) ApiClient {
	cfg := &Config{
		HTTPClient: &http.Client{
			Transport: rt,
		},
		BaseURL: &burl,
	}
	return New(cfg)
}

type mockRoundTripper struct {
	mock.Mock
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	//fmt.Fprintln(GinkgoWriter, actual)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (m *mockRoundTripper) givenResponseBackIs(sc int, body string, err error) {
	hs := make(http.Header)
	hs.Set("Location", "/game/123e4567-e89b-12d3-a456-426655440000")
	res := &http.Response{
		StatusCode: sc,
		// Send response to be tested
		Body: ioutil.NopCloser(bytes.NewBufferString(body)),
		// Must be set to non-nil value or it panics
		Header: hs,
	}
	m.givenResponseReturns(res, err)
}

func (m *mockRoundTripper) givenResponseReturns(res *http.Response, err error) {
	m.On("RoundTrip", mock.AnythingOfType("*http.Request")).
		Return(res, err).
		Once()
}

func (m *mockRoundTripper) assertRequest(url string, method string, obj interface{}) {
	req := m.Calls[0].Arguments[0].(*http.Request)
	Expect(req.URL.String()).To(Equal(url))
	Expect(req.Method).To(Equal(method))
	actual := bytesToString(req.Body)
	expected := jsonToString(obj)
	Expect(actual).Should(MatchJSON(expected))
}

func jsonToString(obj interface{}) string {
	bs, _ := json.Marshal(obj)
	return string(bs)
}

func bytesToString(rc io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	return buf.String()
}
