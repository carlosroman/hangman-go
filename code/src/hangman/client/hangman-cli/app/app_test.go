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
		ta *ApplicationMock
		a  App
	)
	Describe("Application functions", func() {

		BeforeEach(func() {
			ta = &ApplicationMock{}
			a = &app{
				ta: ta,
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

	})
})

type ApplicationMock struct {
	tview.Application
	mock.Mock
}

func (a *ApplicationMock) Run() error {
	fmt.Fprintln(GinkgoWriter, "Run called.")
	return a.Called().Error(0)
}

func (a *ApplicationMock) Stop() {
	fmt.Fprintln(GinkgoWriter, "Stop called.")
	a.Called()
}

func (a *ApplicationMock) SetRoot(root tview.Primitive, fullscreen bool) *tview.Application {
	a.Called(root, fullscreen)
	var at = unsafe.Pointer(a)
	return (*tview.Application)(at)
}

func (a *ApplicationMock) GetWord(d domain.Difficulty) (string, error) {
	args := a.Called(d)
	return args.String(0), args.Error(1)
}
