package sortingandsearching

type sasc struct {
}

func NewSASC() sasc {
	return sasc{}
}

// FindTwoNumbersAddToK finds two numbers in an array that sum a given number
func (sasc) FindTwoNumbersAddToK(arr []int, k int) (int, int) {
	memo := make(map[int]bool)

	for _, el := range arr {
		_, ok := memo[k-el]
		if ok {
			return el, k - el
		}
		memo[el] = true
	}

	return 0, 0
}