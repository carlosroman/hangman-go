package rest

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
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
		cfg := &Config{
			HTTPClient: &http.Client{
				Transport: rt,
			},
			BaseURL: &burl,
		}
		c = New(cfg)
	})

	Describe("New game", func() {

		It("should post correct data", func() {
			rt.givenResponseBackIs(200, "test", nil)
			_, err := c.NewGame("Bob", "easy")
			rt.AssertExpectations(GinkgoT())
			Expect(err).Should(Succeed())
			rt.assertRequest("http://some.base/url/game", "POST")
		})
	})
})

type mockRoundTripper struct {
	mock.Mock
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (m *mockRoundTripper) givenResponseBackIs(sc int, body string, err error) {
	res := &http.Response{
		StatusCode: sc,
		// Send response to be tested
		Body: ioutil.NopCloser(bytes.NewBufferString(body)),
		// Must be set to non-nil value or it panics
		Header: make(http.Header),
	}
	m.givenResponseReturns(res, err)
}

func (m *mockRoundTripper) givenResponseReturns(res *http.Response, err error) {
	m.On("RoundTrip", mock.AnythingOfType("*http.Request")).
		Return(res, err).
		Once()
}

func (m *mockRoundTripper) assertRequest(url string, method string) {
	req := m.Calls[0].Arguments[0].(*http.Request)
	Expect(req.URL.String()).To(Equal(url))
	Expect(req.Method).To(Equal(method))
}
