package sortingandsearching

type insertionSort struct {
}

func NewInsertionSort() insertionSort {
	return insertionSort{}
}

func (s insertionSort) Sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		s.insert(arr, i-1, arr[i])
	}

	return arr
}

func (insertionSort) insert(arr []int, rightIndex int, val int) {
	var i int
	for i = rightIndex; i >= 0 && arr[i] > val; i-- {
		arr[i+1] = arr[i]
	}
	arr[i+1] = val
}
