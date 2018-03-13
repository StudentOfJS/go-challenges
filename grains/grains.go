package grains

import (
	"errors"
	"math"
)

// Square returns the amount of grains on a given square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return uint64(n), errors.New("out of range")
	}
	t := float64(n - 1)
	p := math.Pow(2.0, t)
	return uint64(p), nil
}

// Total returns the total amount of grains if the chessboard is full
func Total() uint64 {
	var total uint64
	for i := 1; i < 65; i++ {
		s, _ := Square(i)
		total += s
	}
	return total
}
