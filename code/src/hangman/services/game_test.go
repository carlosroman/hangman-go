package services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"fmt"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	"hangman/domain"
	"hangman/services"
	"hangman/services/wordstore"
)

var _ = Describe("Game", func() {
	var (
		ws *wordstore.StoreMock
		gs services.GameService
	)

	BeforeEach(func() {
		w := new(wordstore.StoreMock)
		ws = w
		gs = services.NewGameService(ws)
	})

	Describe("Game service", func() {

		Context("on new Game", func() {

			It("should init game correctly", func() {
				ws.On("GetWord", mock.AnythingOfType("domain.Difficulty")).
					Return("word", nil).Once()

				got := gs.NewGame(domain.NORMAL)
				u, err := uuid.FromString(got)
				Expect(err).To(Succeed())
				Expect(u.Version()).To(Equal(uuid.V4))
				Expect(ws.
					AssertCalled(GinkgoT(), "GetWord", domain.NORMAL)).
					To(BeTrue())

			})

			It("should generate unique game ID", func() {
				ws.On("GetWord", mock.AnythingOfType("domain.Difficulty")).
					Return("word", nil).Twice()
				one := gs.NewGame(domain.EASY)
				two := gs.NewGame(domain.HARD)
				Expect(one).ToNot(Equal(two))
				ws.AssertExpectations(GinkgoT())
			})
		})

		Describe("on Get Game", func() {
			BeforeEach(func() {
				ws.On("GetWord", mock.AnythingOfType("domain.Difficulty")).
					Return("word", nil).Once()

			})

			DescribeTable("give the following guesses", func(d domain.Difficulty) {
				id := gs.NewGame(d)
				g := gs.GetGame(id)
				Expect(id).To(Equal(g.Id), "Game IDs should match")
				ws.AssertExpectations(GinkgoT())
			}, Entry("Normal", domain.VERY_EASY),
				Entry("Normal", domain.VERY_HARD),
			)
		})

		Describe("given guessing a letter", func() {

			var (
				id          string
				wordToGuess string
			)
			BeforeEach(func() {
				// letterGuessed: []rune{'w', 'i', 't', 'h', 'd', 'r', 'a', 'w'}
				wordToGuess = "withdraw"
				ws.On("GetWord", mock.AnythingOfType("domain.Difficulty")).
					Return(wordToGuess, nil).
					Once()
				id = gs.NewGame(domain.VERY_HARD)
			})

			type testGuess struct {
				guess         rune
				correct       bool
				missesLeft    int
				gameOver      bool
				letterGuessed []rune
			}

			type guesses []testGuess

			DescribeTable("give the following guesses", func(gss guesses) {
				Expect(gss).NotTo(BeEmpty())
				for _, g := range gss {
					ex, lft, gover, lg := gs.Guess(id, g.guess)
					Expect(ex).To(Equal(g.correct), fmt.Sprintf("Expect guess to be '%t'", g.correct))
					Expect(lft).To(Equal(g.missesLeft), fmt.Sprintf("Expect misses left to be '%d'", g.missesLeft))
					Expect(gover).To(Equal(g.gameOver), fmt.Sprintf("Expect game over to be '%t'", g.gameOver))
					Expect(lg).To(Equal(g.letterGuessed), fmt.Sprintf("Expected '%s' but got '%s'", string(g.letterGuessed), string(lg)))
				}

			}, Entry("First guess", guesses{{
				guess: 'w', correct: true, missesLeft: 8, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
			}), Entry("Two guess, first good, second fail", guesses{
				{guess: 'w', correct: true, missesLeft: 8, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
				{guess: 'b', correct: false, missesLeft: 7, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
			}), Entry("Two guess, first fail, second good", guesses{
				{guess: 'b', correct: false, missesLeft: 7, gameOver: false, letterGuessed: []rune{'_', '_', '_', '_', '_', '_', '_', '_'}},
				{guess: 'w', correct: true, missesLeft: 7, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
			}), Entry("Three guess, 1st good, 2nd bad, 3d good", guesses{
				{guess: 'w', correct: true, missesLeft: 8, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
				{guess: 'b', correct: false, missesLeft: 7, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
				{guess: 'd', correct: true, missesLeft: 7, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', 'd', '_', '_', 'w'}},
			}), Entry("Eight bad guesses", guesses{
				{guess: 'b', correct: false, missesLeft: 7, gameOver: false, letterGuessed: []rune{'_', '_', '_', '_', '_', '_', '_', '_'}},
				{guess: 'b', correct: false, missesLeft: 6, gameOver: false, letterGuessed: []rune{'_', '_', '_', '_', '_', '_', '_', '_'}},
				{guess: 'b', correct: false, missesLeft: 5, gameOver: false, letterGuessed: []rune{'_', '_', '_', '_', '_', '_', '_', '_'}},
				{guess: 'b', correct: false, missesLeft: 4, gameOver: false, letterGuessed: []rune{'_', '_', '_', '_', '_', '_', '_', '_'}},
				{guess: 'b', correct: false, missesLeft: 3, gameOver: false, letterGuessed: []rune{'_', '_', '_', '_', '_', '_', '_', '_'}},
				{guess: 'b', correct: false, missesLeft: 2, gameOver: false, letterGuessed: []rune{'_', '_', '_', '_', '_', '_', '_', '_'}},
				{guess: 'b', correct: false, missesLeft: 1, gameOver: false, letterGuessed: []rune{'_', '_', '_', '_', '_', '_', '_', '_'}},
				{guess: 'b', correct: false, missesLeft: 0, gameOver: true, letterGuessed: []rune{'_', '_', '_', '_', '_', '_', '_', '_'}},
			}), Entry("Eight bad guesses and a correct one still means game over and guess correct is false", guesses{
				{guess: 'w', correct: true, missesLeft: 8, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
				{guess: 'b', correct: false, missesLeft: 7, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
				{guess: 'b', correct: false, missesLeft: 6, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
				{guess: 'b', correct: false, missesLeft: 5, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
				{guess: 'b', correct: false, missesLeft: 4, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
				{guess: 'b', correct: false, missesLeft: 3, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
				{guess: 'b', correct: false, missesLeft: 2, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
				{guess: 'b', correct: false, missesLeft: 1, gameOver: false, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
				{guess: 'b', correct: false, missesLeft: 0, gameOver: true, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
				{guess: 'w', correct: false, missesLeft: 0, gameOver: true, letterGuessed: []rune{'w', '_', '_', '_', '_', '_', '_', 'w'}},
			}),
			)
		})
	})
})
