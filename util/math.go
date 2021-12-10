package util

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func IntPow(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}
