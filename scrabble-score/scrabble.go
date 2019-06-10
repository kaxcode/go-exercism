package scrabble

import (
	"strings"
	"fmt"
)

// Score computes the scrabble score for the word given
func Score(word string) int {
	total := 0
	w := strings.ToUpper(word)
	wordSplit := strings.Split(w, "")

	for _, letter := range wordSplit {
		count := 0
		l := string(letter)
		fmt.Println(l)

		if strings.ContainsAny(l, "A E I O U L N R S T") {
			count++
		}

		if strings.ContainsAny(l, "D G") {
			return count + 2
		}

		if strings.ContainsAny(l, "B C M P") {
			return count + 3
		}

		if strings.ContainsAny(l, "F H V W Y") {
			return count + 4
		}

		if strings.ContainsAny(l, "K") {
			return count + 5
		}

		if strings.ContainsAny(l, "J X") {
			return count + 8
		}

		if strings.ContainsAny(l, "Q Z") {
			return count + 10
		}

		fmt.Println(count)
		return count
	}

	return total
}
