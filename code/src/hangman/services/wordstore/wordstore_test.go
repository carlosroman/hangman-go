package wordstore_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"hangman/domain"
	. "hangman/services/wordstore"
)

var _ = Describe("Wordstore", func() {
	Describe("loading from CSV", func() {
		ws, err := NewInMemoryStoreFromCSV("testdata/simple.csv")
		It("should not error loading CSV", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when CSV loaded correctly", func() {
			It("should return the correct word for VERY_EASY", func() {
				w := ""
				w, err = ws.GetWord(domain.VERY_EASY)
				Expect(err).NotTo(HaveOccurred())
				Expect(w).To(Equal("veryeasy"))
			})
			It("should return the correct word for EASY", func() {
				w := ""
				w, err = ws.GetWord(domain.EASY)
				Expect(err).NotTo(HaveOccurred())
				Expect(w).To(Equal("easy"))
			})
			It("should return the correct word for NORMAL", func() {
				w := ""
				w, err = ws.GetWord(domain.NORMAL)
				Expect(err).NotTo(HaveOccurred())
				Expect(w).To(Equal("normal"))
			})
			It("should return the correct word for HARD", func() {
				w := ""
				w, err = ws.GetWord(domain.HARD)
				Expect(err).NotTo(HaveOccurred())
				Expect(w).To(Equal("hard"))
			})
			It("should return the correct word for VERY_HARD", func() {
				w := ""
				w, err = ws.GetWord(domain.VERY_HARD)
				Expect(err).NotTo(HaveOccurred())
				Expect(w).To(Equal("veryhard"))
			})
		})
	})
})
