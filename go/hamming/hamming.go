package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	length := len(a)
	if length != len(b) {
		return 0, errors.New("Strings are not the same length")
	}

	var mismatches int
	for i := range a {
		if a[i] != b[i] {
			mismatches++
		}
	}

	return mismatches, nil
}
