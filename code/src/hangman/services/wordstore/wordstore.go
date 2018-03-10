package wordstore

import (
	"encoding/csv"
	"errors"
	"fmt"
	"hangman/domain"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	log "hangman/utils"
	"github.com/sirupsen/logrus"
)

var logger = log.Logger()

type Store interface {
	GetWord(d domain.Difficulty) (string, error)
}

type inMemoryStore struct {
	words map[domain.Difficulty][]string
}

func (s *inMemoryStore) GetWord(d domain.Difficulty) (string, error) {
	logger.WithFields(logrus.Fields{
		"difficulty" : d,
	}).Info("Got request for word")
	l := len(s.words[d])
	if l < 1 {
		return "", errors.New("no words available")
	}
	i := rand.Intn(l)
	return s.words[d][i], nil
}

func NewInMemoryStoreFromCSV(csvPath string) (Store, error) {
	fs, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}

	//words := make(map[domain.Difficulty][]string)
	r := csv.NewReader(fs)

	idx := 0
	m := &inMemoryStore{make(map[domain.Difficulty][]string)}
	for {
		record, err := r.Read()
		if err == io.EOF {
			err = nil
			break
		}

		if err != nil {
			break
		}

		if idx > 0 {
			i, err := strconv.Atoi(strings.TrimSpace(record[1]))
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				break
			}

			switch d := domain.Difficulty(i); d {

			case domain.VERY_EASY:
				m.words[domain.VERY_EASY] = append(m.words[domain.VERY_EASY], strings.TrimSpace(record[0]))
			case domain.EASY:
				m.words[domain.EASY] = append(m.words[domain.EASY], strings.TrimSpace(record[0]))
			case domain.NORMAL:
				m.words[domain.NORMAL] = append(m.words[domain.NORMAL], strings.TrimSpace(record[0]))
			case domain.HARD:
				m.words[domain.HARD] = append(m.words[domain.HARD], strings.TrimSpace(record[0]))
			case domain.VERY_HARD:
				m.words[domain.VERY_HARD] = append(m.words[domain.VERY_HARD], strings.TrimSpace(record[0]))
			default:
				err = errors.New(fmt.Sprintf("got Difficulty '%s'", d))
				break
			}
		}
		idx += 1
	}
	return m, err
}
