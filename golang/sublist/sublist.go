package sublist

type Relation = string

func checkSubset(i int, a, b []int) bool {
	lb := len(b)
	for j := 0; j < lb; j++ {
		if a[i] != b[j] {
			return false
		}
		i += 1
	}
	return true
}

func Sublist(a, b []int) string {
	la := len(a)
	lb := len(b)

	if la == lb {
		if checkSubset(0, a, b) {
			return "equal"
		}
	} else if la < lb {
		for i := 0; i < lb; i++ {
			if lb-i >= la && checkSubset(i, b, a) {
				return "sublist"
			}
		}
	} else {
		for i := 0; i < la; i++ {
			if la-i >= lb && checkSubset(i, a, b) {
				return "superlist"
			}
		}
	}
	return "unequal"
}
