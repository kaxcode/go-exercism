// Package scrabble returns scrabble score totals.
package scrabble

import (
	"strings"
)

var scores = map[string]int{}

func init() {
	scores["A"] = 1
	scores["E"] = 1
	scores["I"] = 1
	scores["O"] = 1
	scores["U"] = 1
	scores["L"] = 1
	scores["N"] = 1
	scores["R"] = 1
	scores["S"] = 1
	scores["T"] = 1
	scores["D"] = 2
	scores["G"] = 2
	scores["B"] = 3
	scores["C"] = 3
	scores["M"] = 3
	scores["P"] = 3
	scores["F"] = 4
	scores["H"] = 4
	scores["V"] = 4
	scores["W"] = 4
	scores["Y"] = 4
	scores["K"] = 5
	scores["J"] = 8
	scores["X"] = 8
	scores["Q"] = 10
	scores["Z"] = 10
}

// Score computes the scrabble score for the word given.
func Score(word string) int {
	total := 0
	w := strings.ToUpper(word)

	for _, letter := range w {
		l := string(letter)
		total += scores[l]
	}

	return total
}
