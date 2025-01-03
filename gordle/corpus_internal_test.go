package gordle

import "testing"

func inCorpus(corpus []string, word string) bool {
	for _, corpusWord := range corpus {
		if corpusWord == word {
			return true
		}
	}
	return false
}

func TestPickWord(t *testing.T) {
	corpus := []string{"HELLO", "SALUT", "NUMBER", "XEIFING"}
	word := pickWord(corpus)
	if !inCorpus(corpus, word) {
		t.Errorf("expected a word from the corpus list, got %q", word)
	}
	
}
