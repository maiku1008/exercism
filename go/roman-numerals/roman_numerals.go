package romannumerals

import (
	"fmt"
	"strings"
)

type value struct {
	arabic int
	roman  string
}

var numerals = []value{
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
	for _, v := range numerals {
		for num >= v.arabic {
			result.WriteString(v.roman)
			num -= v.arabic
		}
	}
	return result.String(), nil
}
