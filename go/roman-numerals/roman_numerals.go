package romannumerals

import (
	"fmt"
	"strings"
)

type numeral struct {
	arabic int
	roman  string
}

var numerals = []numeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(num int) (string, error) {
	if num <= 0 || num > 3000 {
		return "", fmt.Errorf("number should be between 0 and 3000: %d", num)
	}
	var result strings.Builder
	// 	for num > 0 {
	// 		switch {
	// 		case num >= 1000:
	// 			result.WriteString("M")
	// 			num -= 1000
	// 		case num >= 900:
	// 			result.WriteString("CM")
	// 			num -= 900
	// 		case num >= 500:
	// 			result.WriteString("D")
	// 			num -= 500
	// 		case num >= 400:
	// 			result.WriteString("CD")
	// 			num -= 400
	// 		case num >= 100:
	// 			result.WriteString("C")
	// 			num -= 100
	// 		case num >= 90:
	// 			result.WriteString("XC")
	// 			num -= 90
	// 		case num >= 50:
	// 			result.WriteString("L")
	// 			num -= 50
	// 		case num >= 40:
	// 			result.WriteString("XL")
	// 			num -= 40
	// 		case num >= 10:
	// 			result.WriteString("X")
	// 			num -= 10
	// 		case num >= 9:
	// 			result.WriteString("IX")
	// 			num -= 9
	// 		case num >= 5:
	// 			result.WriteString("V")
	// 			num -= 5
	// 		case num >= 4:
	// 			result.WriteString("IV")
	// 			num -= 4
	// 		case num >= 1:
	// 			result.WriteString("I")
	// 			num -= 1
	// 		}
	// 	}
	for _, v := range numerals {
		for num >= v.arabic {
			result.WriteString(v.roman)
			num -= v.arabic
		}
	}
	return result.String(), nil
}
