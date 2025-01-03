package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

func New(playerInput io.Reader, corpus []string, maxAttempts int) (*Game, error) {
	if len(corpus) == 0 {
		return nil, ErrCorpusIsEmpty
	}
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    toUpperChars(pickWord(corpus)),
		maxAttempts: maxAttempts,
	}

	return g, nil
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle")
	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()

		fb := computeFeedback(guess, g.solution)
		fmt.Println(fb.String())

		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉 You won! You found it in %d guess(es)\n", currentAttempt)
			return
		}
	}
	fmt.Printf("🥹  You lost! The solution was %s\n", string(g.solution))
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess: ", len(g.solution))

	playerInput, _, err := g.reader.ReadLine()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read")
		return nil
	}
	guess := toUpperChars(string(playerInput))
	err = g.validateGuess(guess)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s\n", err.Error())
		return nil
	}
	return guess
}

var errInvalidWordLength = fmt.Errorf("guess doesn't have the same number of characters as the solution")

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}
	return nil
}

func toUpperChars(word string) []rune {
	return []rune(strings.ToUpper(word))
}
