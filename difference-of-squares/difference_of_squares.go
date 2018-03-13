package diffsquares

import "math"

// SquareOfSums sums all natural numbers up to provided int and then returns the square
func SquareOfSums(squares int) int {
	var sum float64
	for i := 1; i <= squares; i++ {
		sum += float64(i)
	}
	return int(math.Pow(sum, 2.0))
}

// SumOfSquares squares all natural numbers up to provided int and then returns the sum
func SumOfSquares(squares int) int {
	var sum float64
	for i := 1; i <= squares; i++ {
		sum += math.Pow(float64(i), 2.0)
	}
	return int(sum)
}

// Difference returns the difference between the SquareOfSums and SumOfSquares
func Difference(squares int) int {
	return SquareOfSums(squares) - SumOfSquares(squares)
}
