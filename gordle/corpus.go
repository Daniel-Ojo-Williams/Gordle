package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type corpusError string

func (e corpusError) Error() string {
	return string(e)
}

const ErrCorpusIsEmpty = corpusError("corpus is empty")

func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading: %w", path, ErrCorpusIsEmpty)
	}
	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	words := strings.Fields(string(data))
	return words, nil
}

func pickWord(corpus []string) string {
	index := rand.Intn(len(corpus))
	return corpus[index]
}