package sortingandsearching

type selectionSort struct {
}

func NewSelectionSort() selectionSort {
	return selectionSort{}
}

func (s selectionSort) Sort(arr []int) []int {
	minIndex := 0
	for i := 0; i < len(arr)-1; i++ {
		minIndex = s.indexOfMinimum(arr, i)
		s.swap(arr, i, minIndex)
	}

	return arr
}

func (selectionSort) indexOfMinimum(arr []int, startIndex int) int {
	minValue := arr[startIndex]
	minIndex := startIndex

	for i := minIndex + 1; i < len(arr); i++ {
		if arr[i] < minValue {
			minIndex = i
			minValue = arr[i]
		}
	}

	return minIndex
}

func (selectionSort) swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
