package util

func RuneAt(str string, index int) rune {
	for i, r := range str {
		if i == index {
			return r
		}
	}
	return 0
}
