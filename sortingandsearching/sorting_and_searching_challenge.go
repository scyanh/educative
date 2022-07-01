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

// FindPivotIndex finds pivot index in a rotated array sorted
func (sasc) FindPivotIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 {
		return 0
	}

	if arr[0] < arr[len(arr)-1] {
		return 0
	}

	left, right := 0, len(arr)-1

	for left < right {
		mid := (left + right) / 2

		if arr[mid] > arr[mid+1] {
			return mid + 1
		}

		if arr[mid] < arr[mid-1] {
			return mid
		}

		if arr[mid] > arr[0] {
			left = mid
		} else {
			right = mid
		}
	}

	return 0
}