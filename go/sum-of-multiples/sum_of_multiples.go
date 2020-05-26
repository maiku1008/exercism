package summultiples

// SumMultiples sums the multiples of up to and including limit
func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for num := 1; num < limit; num++ {
		for _, divisor := range divisors {
			if divisor != 0 && num%divisor == 0 {
				sum += num
				break
			}
		}
	}
	return sum
}
