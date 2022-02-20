package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	if i == nil {
		return nil
	}
	res := []int{}
	for _, number := range i {
		if filter(number) {
			res = append(res, number)
		}
	}
	return res
}

func (i Ints) Discard(filter func(int) bool) Ints {
	return i.Keep(func(n int) bool { return !filter(n) })
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	if l == nil {
		return nil
	}
	res := [][]int{}
	for _, list := range l {
		if filter(list) {
			res = append(res, list)
		}
	}
	return res
}

func (s Strings) Keep(filter func(string) bool) Strings {
	if s == nil {
		return nil
	}
	res := []string{}
	for _, word := range s {
		if filter(word) {
			res = append(res, word)
		}
	}
	return res
}
