package diffsquares

// SquareOfSum calculates the square of the sum of the first n integers
func SquareOfSum(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares calculates the sum of the first n int integers
func SumOfSquares(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

// Difference calculates the difference of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
