package main_test

import (
	"github.com/gdamore/tcell"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rivo/tview"

	. "hangman/client/hangman-cli"
)

var _ = Describe("Start", func() {
	var (
		startScreen tview.Primitive
		quit        bool
	)

	BeforeEach(func() {
		quit = false
		startScreen = Start(func(p string) {}, func() {
			quit = true
		})
	})

	Describe("Quitting works", func() {
		Context("Clicking quit button", func() {
			It("should call stop function", func() {
				k := tcell.NewEventKey(tcell.KeyTAB, ' ', tcell.ModNone)
				f := startScreen.(*tview.Form)
				var bp tview.Primitive
				f.Focus(func(p tview.Primitive) {
					bp = p
				})
				f.GetFormItem(0).InputHandler()(k, func(p tview.Primitive) {})
				bp.InputHandler()(k, func(p tview.Primitive) {})
				bp.InputHandler()(k, func(p tview.Primitive) {})
				Expect(bp).NotTo(BeNil())
				bb := bp.(*tview.Button)
				bb.InputHandler()(tcell.NewEventKey(tcell.KeyEnter, ' ', tcell.ModNone), func(p tview.Primitive) {})
				Expect(quit).To(BeTrue())
			})
		})
	})
})
