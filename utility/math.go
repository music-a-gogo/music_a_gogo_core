package utility

// GCD represents the Greatest Common Divisor, using Euclidean algorithm.
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM represents the Least Common Multiple, using Euclid's algorithm.
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func Sum(a []int) int {
	sum := 0
	for _, n := range a {
		sum += n
	}
	return sum
}
