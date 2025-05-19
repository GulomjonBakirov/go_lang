package syntax

func AddSum(a, b int) int {
	return a + b
}

func Substract(a, b int) (bool, int) {
	if a <= b {
		return false, 0
	}

	return true, a - b
}
