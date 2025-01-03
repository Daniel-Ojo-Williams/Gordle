package main

import (
	"fmt"
	"os"

	"github.com/daniel-ojo-williams/gordle/gordle"
)

func main() {
	corpus, err := gordle.ReadCorpus("./corpus/english.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to read corpus file: %s", err.Error())
		return
	}
	game, err := gordle.New(os.Stdin, corpus, 6)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to start game: %s", err.Error())
		return
	}
	game.Play()
}