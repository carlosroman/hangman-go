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
		f           *tview.Form
		bp          tview.Primitive
		nickname    string
		difficulty  string
	)

	Describe("Start Screen loads", func() {

		BeforeEach(func() {
			quit = false
			startScreen = Start(func(p string) {}, func() {
				quit = true
			}, func(n string, d string) error {
				nickname = n
				difficulty = d
				return nil
			})
			f = startScreen.(*tview.Form)
			f.Focus(func(p tview.Primitive) {
				bp = p
			})
		})

		Context("when the  player wants to quit", func() {

			It("should call stop function on clicking quit", func() {
				k := tcell.NewEventKey(tcell.KeyTAB, ' ', tcell.ModNone)
				f.GetFormItem(0).InputHandler()(k, func(p tview.Primitive) {})
				bp.InputHandler()(k, func(p tview.Primitive) {})
				bp.InputHandler()(k, func(p tview.Primitive) {})
				Expect(bp).NotTo(BeNil())
				bb := bp.(*tview.Button)
				bb.InputHandler()(tcell.NewEventKey(tcell.KeyEnter, ' ', tcell.ModNone), func(p tview.Primitive) {})
				Expect(quit).To(BeTrue())
			})
		})

		Context("when the player wants to play", func() {

			It("should send correct new game data", func() {
				f.GetFormItem(0).InputHandler()(
					tcell.NewEventKey(tcell.KeyRune, 'B', tcell.ModNone),
					func(p tview.Primitive) {})
				f.GetFormItem(0).InputHandler()(
					tcell.NewEventKey(tcell.KeyRune, 'o', tcell.ModNone),
					func(p tview.Primitive) {})
				f.GetFormItem(0).InputHandler()(
					tcell.NewEventKey(tcell.KeyRune, 'b', tcell.ModNone),
					func(p tview.Primitive) {})
				k := tcell.NewEventKey(tcell.KeyTAB, ' ', tcell.ModNone)
				f.GetFormItem(0).InputHandler()(k, func(p tview.Primitive) {})
				bp.InputHandler()(k, func(p tview.Primitive) {})
				Expect(bp).NotTo(BeNil())
				bb := bp.(*tview.Button)
				bb.InputHandler()(tcell.NewEventKey(tcell.KeyEnter, ' ', tcell.ModNone), func(p tview.Primitive) {})
				Expect(difficulty).To(Equal("Very easy"))
				Expect(nickname).To(Equal("Bob"))
			})
		})
	})
})
