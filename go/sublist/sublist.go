package sublist

// Relation is a string
type Relation string

// Sublist returns a Relation according to the relation of slices a and b
func Sublist(a, b []int) Relation {
	switch {
	case len(a) == len(b) && isSublist(a, b):
		return "equal"
	case len(a) > len(b) && isSublist(b, a):
		return "superlist"
	case len(a) < len(b) && isSublist(a, b):
		return "sublist"
	default:
		return "unequal"
	}
}

func isSublist(small, large []int) bool {
	for i := 0; i < len(large)-len(small)+1; i++ {
		match := true
		for j, v := range small {
			if v != large[i+j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}




