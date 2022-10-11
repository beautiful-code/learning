package etl

func Transform(in map[int][]string) map[string]int {
	result := map[string]int{}
	helper := map[string]string{"A": "a", "B": "b", "C": "c", "D": "d", "E": "e", "F": "f", "G": "g", "H": "h", "I": "i", "J": "j", "K": "k", "L": "l", "M": "m", "N": "n", "O": "o", "P": "p", "Q": "q", "R": "r", "S": "s", "T": "t", "U": "u", "V": "v", "W": "w", "X": "x", "Y": "y", "Z": "z"}
	for key, val := range in {
		for _, letter := range val {
			result[helper[letter]] = key
		}
	}
	return result
}
