package scrabble

import (
	"strings"
)

// Score computes the scrabble score for the word given
func Score(word string) int {
	total := 0
	w := strings.ToUpper(word)

	for _, letter := range w {
		l := string(letter)
		if strings.ContainsAny(l, "A E I O U L N R S T") {
			total++
		} else if strings.ContainsAny(l, "D G") {
			total += 2
		} else if strings.ContainsAny(l, "B C M P") {
			total += 3
		} else if strings.ContainsAny(l, "F H V W Y") {
			total += 4
		} else if strings.ContainsAny(l, "K") {
			total += 5
		} else if strings.ContainsAny(l, "J X") {
			total += 8
		} else if strings.ContainsAny(l, "Q Z") {
			total += 10
		}
	}

	return total
}
