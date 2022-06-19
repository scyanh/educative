package recursion

type recursionPower struct {
}

func NewRecursionPower() recursionPower {
	return recursionPower{}
}

// Power returns the number power receiving a number (x) and a power (n)
// example x=2 and n=3 returns 8
// 2 * 2 * 2 = 8
func (r recursionPower) Power(x, n int) int {
	if n == 0 {
		return 1
	}

	return x * r.Power(x, n-1)
}
