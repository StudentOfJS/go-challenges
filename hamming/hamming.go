package hamming

import (
	"errors"
)

// Distance measures the hamming difference between two strands of DNA
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		err := errors.New("Not equal")
		return -1, err
	}
	diffs := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		diffs++
	}
	return diffs, nil
}
