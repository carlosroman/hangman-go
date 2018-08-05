package wordstore_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"hangman/domain"
	. "hangman/services/wordstore"
)

var _ = Describe("Wordstore", func() {
	var (
		ws  Store
		err error
	)

	Describe("loading from CSV", func() {
		BeforeEach(func() {
			ws, err = NewInMemoryStoreFromCSV("testdata/simple.csv")
			Expect(err).To(Succeed())
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

	Describe("loading from CSV can fail", func() {

		It("should fail if can't find file", func() {
			ws, err = NewInMemoryStoreFromCSV("bad/file/path/to.csv")
			Expect(err).NotTo(Succeed())
			Expect(ws).To(BeNil())
		})

		It("should fail if can't match difficulty", func() {
			ws, err = NewInMemoryStoreFromCSV("testdata/bad_difficulty.csv")
			Expect(err).NotTo(Succeed())
		})

		It("should fail if wrong number of records on a line", func() {
			ws, err = NewInMemoryStoreFromCSV("testdata/incorrec_records.csv")
			Expect(err).NotTo(Succeed())
		})

		It("should fail if CSV file bad", func() {
			ws, err = NewInMemoryStoreFromCSV("testdata/bad.csv")
			Expect(err).NotTo(Succeed())
		})
	})
})
