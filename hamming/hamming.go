package hamming

import "errors"

// Distance returns a count of differences between two homologous DNA strands
func Distance(a, b string) (int, error) {
	count := 0

	if len(a) != len(b) {
		return count, errors.New("Sequences are not of equal length")
	}

	for i := range a {
		if a[i] == b[i] {
			continue
		}
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil

}
