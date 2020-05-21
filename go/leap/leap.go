package leap

func IsLeapYear(year int) bool {
	switch {
	case (year%4 == 0) && (year%100 == 0) && (year%400 == 0):
		return true
	case (year%4 == 0) && (year%100 == 0):
		return false
	case (year%4 == 0):
		return true
	default:
		return false
	}
}
