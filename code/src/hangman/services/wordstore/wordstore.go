package wordstore

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"hangman/domain"
	log "hangman/utils"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
		"difficulty": d,
	}).Info("Got request for word")
	l := len(s.words[d])
	if l < 1 {
		return "", errors.New("no words available")
	}
	i := rand.Intn(l)
	return s.words[d][i], nil
}

func NewInMemoryStoreFromCSV(csvPath string) (ws Store, err error) {
	fs, err := os.Open(csvPath)
	if err != nil {
		return ws, err
	}

	//words := make(map[domain.Difficulty][]string)
	r := csv.NewReader(fs)

	idx := 0
	m := &inMemoryStore{
		words: make(map[domain.Difficulty][]string),
	}

	for {

		record, rerr := r.Read()

		if rerr != nil {
			if rerr == io.EOF {
				ws = m
				logger.WithFields(logrus.Fields{
					"csvPath": csvPath,
				}).Infoln("Done Loading CSV")
				break
			}
			err = rerr
			logger.WithFields(logrus.Fields{
				"csvPath": csvPath,
			}).Error(err)
			break
		}

		if idx > 0 {
			if len(record) != 2 {
				err = fmt.Errorf("found '%d' records", len(record))
				logger.WithFields(logrus.Fields{
					"csvPath": csvPath,
				}).Error(err)
				break
			}

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
				err = fmt.Errorf("got Difficulty '%s'", d)
				logger.WithFields(logrus.Fields{
					"csvPath": csvPath,
				}).Error(err)
				return m, err
			}
		}
		idx += 1
	}
	return ws, err
}
