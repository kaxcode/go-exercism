package twofer

// ShareWith returns a string
func ShareWith(name string) string {
	if len(name) != 0 {
		return "One for " + name + ", one for me."
	}

	return "One for you, one for me."
}
