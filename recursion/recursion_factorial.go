package recursion

type recursiveFactorial struct {
}

func NewRecursiveFactorial() recursiveFactorial {
	return recursiveFactorial{}
}

// Factorial (n)
// example n=5
// 5 * 4 * 3 * 2 * 1 = 120
func (r recursiveFactorial) Factorial(i int) int {
	if i == 0 {
		return 1
	}

	return i * r.Factorial(i-1)
}
