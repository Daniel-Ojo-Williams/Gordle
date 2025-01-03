package gordle

import "strings"

type hint byte

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "🟫"
	case wrongPosition:
		return "🟡"
	case correctPosition:
		return "💚"
	default:
		return "💔"
	}
}

type feedback []hint

func (f feedback) String() string {
	sb := strings.Builder{}
	for _, h := range f {
		sb.WriteString(h.String())
	}
	return sb.String()
}

func computeFeedback(guess, solution []rune) feedback {
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	for positionInGuess, character := range guess {
		if character == solution[positionInGuess] {
			result[positionInGuess] = correctPosition
			used[positionInGuess] = true
		}
	}

	for positionInGuess, character := range guess {
		if result[positionInGuess] != absentCharacter {
			continue
		}
		for positionInSolution, target := range solution {
			if used[positionInSolution] {
				continue
			}
			if character == target {
				result[positionInGuess] = wrongPosition
				used[positionInSolution] = true
				break
			}
		}
	}

	return result
}

func (f feedback) Equal(other feedback) bool {
	if len(f) != len(other) {
		return false
	}

	for index, value := range f {
		if value != other[index] {
			return false
		}
	}

	return true
}
