package app

import (
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hangman/domain"
	"unsafe"
)

var _ = Describe("App", func() {
	var (
		ta *TvViewApplicationMock
		ac *ApiClientMock
		a  App
	)
	Describe("Application functions", func() {

		BeforeEach(func() {
			ta = &TvViewApplicationMock{}
			ac = &ApiClientMock{}
			a = &app{
				ta: ta,
				ac: ac,
			}
		})

		Context("when stop called", func() {

			It("should stop the application", func() {
				ta.On("Stop").Return()
				a.StopApp()
				Expect(true).To(BeTrue())
				assert.True(
					GinkgoT(),
					ta.AssertCalled(GinkgoT(), "Stop"),
					"Expected Stop to be called")
			})
		})

		Context("when run called", func() {

			It("should run the application on success", func() {
				ta.On("Run").Return(nil)
				Expect(a.Run()).Should(Succeed())
			})

			It("should return error if application won't run", func() {
				ta.On("Run").Return(errors.New("failed to run"))
				Expect(a.Run()).ShouldNot(Succeed())
			})
		})

		Context("when new game called", func() {

			It("should call the client correctly", func() {
				ac.givenNewGameReturns("some-id", nil)
				Expect(a.NewGame("name", "difficulty")).Should(Succeed())
				ac.AssertExpectations(GinkgoT())
				assert.True(
					GinkgoT(),
					ac.AssertCalled(GinkgoT(), "NewGame", "name", "difficulty"),
					"Expected New game to be called")
			})

			It("should store the current game id", func() {
				ac.givenNewGameReturns("some-id", nil)
				a.NewGame("name", "difficulty")
				Expect(a.GetGameId()).To(Equal("some-id"))
			})

			It("should return the client error", func() {
				err := errors.New("some error")
				ac.givenNewGameReturns("some-id", err)
				Expect(a.NewGame("name", "difficulty")).ShouldNot(Succeed())
				ac.AssertExpectations(GinkgoT())
			})
		})
	})
})

type TvViewApplicationMock struct {
	tview.Application
	mock.Mock
}

func (a *TvViewApplicationMock) Run() error {
	fmt.Fprintln(GinkgoWriter, "Run called.")
	return a.Called().Error(0)
}

func (a *TvViewApplicationMock) Stop() {
	fmt.Fprintln(GinkgoWriter, "Stop called.")
	a.Called()
}

func (a *TvViewApplicationMock) SetRoot(root tview.Primitive, fullscreen bool) *tview.Application {
	a.Called(root, fullscreen)
	var at = unsafe.Pointer(a)
	return (*tview.Application)(at)
}

func (a *TvViewApplicationMock) GetWord(d domain.Difficulty) (string, error) {
	args := a.Called(d)
	return args.String(0), args.Error(1)
}

type ApiClientMock struct {
	mock.Mock
}

func (c *ApiClientMock) NewGame(name string, difficulty string) (string, error) {
	args := c.Called(name, difficulty)
	return args.String(0), args.Error(1)
}

func (c *ApiClientMock) givenNewGameReturns(id string, err error) {
	c.
		On("NewGame", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(id, err)
}
