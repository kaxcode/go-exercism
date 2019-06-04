package acronym

import (
	"strings"
)

// Abbreviate return string Acronym
func Abbreviate(s string) string {
	// Clean the string and split into a list
	r := strings.NewReplacer("-", " ", "_", "")
	s = r.Replace(s)
	stringReplaced := strings.Replace(s, "  ", "", -1)
	list := strings.SplitAfter(stringReplaced, " ")

	// Iterate over each item and grab the first letter
	newList := []string{}
	for _, word := range list {
		x := strings.ToUpper(word)
		newList = append(newList, string(x[0]))
	}

	// Join list and return string
	result := strings.Join(newList, "")
	return result
}
