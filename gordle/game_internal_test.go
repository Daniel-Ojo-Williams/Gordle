package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in english": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in arabic": {
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		"5 characters in japanesse": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in japanese": {
			input: "ありがとう",
			want:  []rune("ありがとう"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g, err := New(strings.NewReader(tc.input), []string{string(tc.want)}, 0)
			if err != nil {
				t.Errorf("not expecting error but got :%s", err)
			}

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %s, want = %s", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word     []rune
		expected error
	}{
		"too many characters": {
			word:     []rune("WORD OF THE LORD"),
			expected: errInvalidWordLength,
		},
		"correct guess length": {
			word:     []rune("RIGGS"),
			expected: nil,
		},
		"too few characters": {
			word:     []rune("WOR"),
			expected: errInvalidWordLength,
		},
		"nil": {
			word:     nil,
			expected: errInvalidWordLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			game, err := New(nil, []string{"RIGGS"}, 1)
			if err != nil {
				t.Errorf("not expecting an error, got %s", err)
			}
			got := game.validateGuess(tc.word)

			if !errors.Is(got, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestComputeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"nominal": {
			guess:            "PHONE",
			solution:         "PHONE",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character": {
			guess:            "GUESS",
			solution:         "GUESS",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character with wrong answer": {
			guess:            "GUSSS",
			solution:         "GUESS",
			expectedFeedback: feedback{correctPosition, correctPosition, absentCharacter, correctPosition, correctPosition},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !fb.Equal(tc.expectedFeedback) {
				t.Errorf("guess: %q, got the wrong feedback, expected: %v", tc.guess, tc.expectedFeedback)
			}
		})
	}
}
