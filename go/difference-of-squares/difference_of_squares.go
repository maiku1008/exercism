package diffsquares

// SquareOfSum returns the square of the sum of 1..n
func SquareOfSum(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i
	}
	return result * result
}

// SumOfSquares returns the sum of the squares of 1..n
func SumOfSquares(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

// Difference returns the difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	result := (((3 * (n * n)) + (2 * n)) * (1 - (n * n)) / 12)
	if result < 0 {
		result = -result
	}
	return result
}
