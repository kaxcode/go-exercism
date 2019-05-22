package twofer

import "fmt"

// ShareWith returns a string
func ShareWith(name string) string {
	s := "One for %s, one for me."
	if len(name) != 0 {
		return fmt.Sprintf(s, name)
	}

	return fmt.Sprintf(s, "you")
}
