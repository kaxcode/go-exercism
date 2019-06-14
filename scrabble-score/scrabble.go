// Package scrabble returns scrabble score totals.
package scrabble

import (
	"strings"
)

// Score computes the scrabble score for the word given.ß
func Score(word string) int {
	total := 0
	w := strings.ToUpper(word)

	for _, letter := range w {
		l := string(letter)
		switch {
		case strings.ContainsAny(l, "A E I O U L N R S T"):
			total++
		case strings.ContainsAny(l, "D G"):
			total += 2
		case strings.ContainsAny(l, "B C M P"):
			total += 3
		case strings.ContainsAny(l, "F H V W Y"):
			total += 4
		case strings.ContainsAny(l, "K"):
			total += 5
		case strings.ContainsAny(l, "J X"):
			total += 8
		case strings.ContainsAny(l, "Q Z"):
			total += 10
		}
	}

	return total
}
