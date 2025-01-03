package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/daniel-ojo-williams/gordle/gordle"
)

func main() {
	file := flag.String("file", "", "file to read corpus from")
	flag.Parse()
	if *file == "" {
		fmt.Fprint(os.Stderr, "please provide a valid corpus file\n")
		return
	}
	corpus, err := gordle.ReadCorpus(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to read corpus file: %s\n", err.Error())
		return
	}
	game, err := gordle.New(os.Stdin, corpus, 6)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to start game: %s\n", err.Error())
		return
	}
	game.Play()
}