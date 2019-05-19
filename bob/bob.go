package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	yelling := isYelling(remark)
	question := isAsking(remark)
	numeric := isNumericOnly(remark)
	yellingWithNumbers := isYellingWithNumbers(remark)
	numericQuestion := isNumericQuestion(remark)
	noLettersQuestion := isNonLettersQuestions(remark)
	yellingQuestion := isYellingAQuestion(remark)
	silent := isSilent(remark)

	if silent {
		return "Fine. Be that way!"
	} else if numeric {
		return "Whatever."
	} else if numericQuestion || noLettersQuestion {
		return "Sure."
	} else if yellingWithNumbers {
		return "Whoa, chill out!"
	} else if yellingQuestion {
		return "Calm down, I know what I'm doing!"
	} else if question {
		return "Sure."
	} else if yelling {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func isYelling(str string) bool {
	return str == strings.ToUpper(str)
}

func isAsking(str string) bool {
	return strings.HasSuffix(str, "?")
}

func isSilent(str string) bool {
	return !strings.ContainsAny(str, "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ!?") || len(str) == 0
}

func isNumericOnly(str string) bool {
	return (strings.ContainsAny(str, "1234567890") && !strings.ContainsAny(str, "ABCDEFGHIJKLMNOPQRSTUVWXYZ!?"))
}

func isYellingWithNumbers(str string) bool {
	return str == strings.ToUpper(str) && strings.ContainsAny(str, "1234567890")
}

func isYellingAQuestion(str string) bool {
	return str == strings.ToUpper(str) && strings.HasSuffix(str, "?")
}

func isNumericQuestion(str string) bool {
	return strings.ContainsAny(str, "1234567890") && strings.Contains(str, "?")
}

func isNonLettersQuestions(str string) bool {
	return strings.HasSuffix(str, "?") && !strings.ContainsAny(str, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
