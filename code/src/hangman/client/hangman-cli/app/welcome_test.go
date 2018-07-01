package app

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rivo/tview"

	"fmt"
	"github.com/gdamore/tcell"
	"reflect"
)

var _ = Describe("Welcome", func() {

	var (
		welcomeScreen tview.Primitive
		quit          bool
		next          string
		s             tcell.Screen
		f             tview.Primitive
	)

	BeforeEach(func() {
		s = tcell.NewSimulationScreen("")
		Expect(s.Init()).Should(Succeed())
		quit = false
		x, y := s.Size()
		Expect(x).To(Equal(80))
		Expect(y).To(Equal(25))

		welcomeScreen = Welcome(func(p string) {
			next = p
		}, func() {
			quit = true
		})
	})

	AfterEach(func() {
		s.Fini()
	})

	Describe("Welcome screen actions", func() {

		Context("when the screen loads", func() {

			BeforeEach(func() {
				welcomeScreen.Draw(s)
				s.Show()
				welcomeScreen.Focus(func(p tview.Primitive) {
					fmt.Fprintln(GinkgoWriter, reflect.TypeOf(p))
					p.Focus(func(p tview.Primitive) {
						fmt.Fprintln(GinkgoWriter, reflect.TypeOf(p))
						f = p
					})
				})

				Expect(f).NotTo(BeNil())

				f.(*tview.Form).Focus(func(p tview.Primitive) {
					fmt.Fprintln(GinkgoWriter, reflect.TypeOf(p))
					fmt.Fprintln(GinkgoWriter, p.(*tview.Button).GetLabel())
					f = p
				})
			})

			It("should have a start button that works", func() {

				Expect(f.(*tview.Button).GetLabel()).To(Equal("Play"))

				e := tcell.NewEventKey(tcell.KeyEnter, ' ', tcell.ModNone)
				f.(*tview.Button).InputHandler()(e, func(b tview.Primitive) {})

				Expect(next).To(Equal("start"))
			})

			It("should have a quit button that works", func() {

				k := tcell.NewEventKey(tcell.KeyTAB, ' ', tcell.ModNone)
				e := tcell.NewEventKey(tcell.KeyEnter, ' ', tcell.ModNone)

				f.(*tview.Button).InputHandler()(k, func(b tview.Primitive) {})
				Expect(f.(*tview.Button).GetLabel()).To(Equal("Quit"))
				f.(*tview.Button).InputHandler()(e, func(b tview.Primitive) {})

				Expect(quit).To(BeTrue())
			})
		})
	})
})
