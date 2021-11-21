package mathutils

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Multiply(n1 int, n2 int) int {
	return n1 * n2
}

func SafeMe() bool {
	return false
}
