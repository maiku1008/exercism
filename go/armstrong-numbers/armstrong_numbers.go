package armstrong

func IsNumber(number int) bool {
	var sum, size int
	for n := number; n != 0; n /= 10 {
		size++
	}
	for n := number; n != 0; n /= 10 {
		sum += pow(n%10, size)
	}
	return sum == number
}

func pow(n, e int) int {
	if e == 0 {
		return 1
	}
	result := n
	for i := 2; i <= e; i++ {
		result *= n
	}
	return result
}
