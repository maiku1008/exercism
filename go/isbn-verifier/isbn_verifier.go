package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	result := 0
	length := 10
	if len(isbn) != length {
		return false
	}

	for i := 0; i < length; i++ {
		c := string(isbn[i])
		if i == length-1 && c == "X" {
			c = "10"
		}
		digit, err := strconv.Atoi(c)
		if err != nil {
			return false
		}
		result += digit * (length - i)
	}

	return result%11 == 0
}
