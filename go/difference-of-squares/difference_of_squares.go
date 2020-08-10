package diffsquares

// SumOfSquares is the sum of the squares in range [1,n]
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// SquareOfSum is the sum of numbers in range [1,n], squared
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// Difference is the difference betwee sum of squares and square of sums
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
