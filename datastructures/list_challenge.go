package datastructures

import (
	"errors"
)

type ListChallenge struct {
}

type IListChallenge interface {
	RemoveEven(arr []int) []int
	MergeTwoSortedList(arr1, arr2 []int) []int
	TwoSum(arr []int, k int) (int, int, error)
	TwoSum2(arr []int, k int) ([]int, error)
	FindProductAllElements(arr []int) []int
	FindProductAllElements2(arr []int) []int
	FindMinimum(arr []int) int
	FindMinimum2(arr []int) int
	FindSecondMaximum(arr []int) int
	FindSecondMaximum2(arr []int) int
	RearrangeArray(arr []int) []int
	RearrangeArray2(arr []int) []int
	ArrangeMaxMinElements(arr []int) []int
	ArrangeMaxMinElements2(arr []int) []int
	FindMaxSumSublist(arr []int) ([]int, int)
	FindMaxSumSublist2(arr []int) ([]int, int)
}

func NewListChallenge() IListChallenge {
	return ListChallenge{}
}

// arr := []int{1, 2, 3, 4, 5}
// [1 3 5]
func (ListChallenge) RemoveEven(arr []int) []int {
	res := make([]int, 0)
	for _, el := range arr {
		if el%2 != 0 {
			res = append(res, el)
		}
	}

	return res
}

// arr := []int{1, 2, 3, 4, 5}
// arr2 := []int{1, 2, 6, 7}
// [1 1 2 2 3 4 5 6 7]
func (ListChallenge) MergeTwoSortedList(arr1, arr2 []int) []int {
	res := make([]int, 0)

	idx1, idx2 := 0, 0
	total := len(arr1) + len(arr2)
	for i := 0; i < total; i++ {
		// check len arr2 to prevent index out of range
		if idx2 >= len(arr2) || idx1 < len(arr1) && arr1[idx1] <= arr2[idx2] {
			res = append(res, arr1[idx1])
			idx1++
		} else {
			res = append(res, arr2[idx2])
			idx2++
		}
	}

	return res
}

// encontrar dos números en arr cuya suma sea k
// arr := []int{1, 2, 3, 4, 5}
// k := 5
// output [2, 3]
func (ListChallenge) TwoSum(arr []int, k int) (int, int, error) {
	memo := make(map[int]bool)

	for _, el := range arr {
		_, ok := memo[k-el]
		if ok {
			return el, k - el, nil
		}
		memo[el] = true
	}

	return 0, 0, errors.New("not found")
}
func (ListChallenge) TwoSum2(arr []int, k int) ([]int, error) {
	dict := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		diff := k - arr[i]
		if _, ok := dict[diff]; ok {
			return []int{arr[i], diff}, nil
		} else {
			dict[arr[i]] = i
		}
	}

	return []int{}, nil

}

// FindProductAllElements List of Products of all Elements
/*
Sample Input [1,2,3,4]
output = [24,12,8,6]
1=     2 * 3 * 4
2= 1     * 3 * 4
3= 1 * 2     * 4 // i=2
4= 1 * 2 * 3

Sample Input [0, 1, 2, 3]
output [6, 0, 0, 0]
*/
func (ListChallenge) FindProductAllElements(arr []int) []int {
	res := make([]int, len(arr))

	// get product start from left
	left := 1
	for i, el := range arr {
		res[i] = left
		left = left * el
	}

	// get product starting from right
	right := 1
	for i := len(arr) - 1; i >= 0; i-- {
		res[i] = res[i] * right
		right = right * arr[i]
	}

	return res
}
func (ListChallenge) FindProductAllElements2(arr []int) []int {
	left := 1
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = left
		left *= arr[i]
	}

	right := 1
	for i := len(arr) - 1; i >= 0; i-- {
		res[i] *= right
		right *= arr[i]
	}

	return res
}

// FindMinimum finds the smallest number in the given list.
// Sample Input [5, 2, 3]
// output 2
func (ListChallenge) FindMinimum(arr []int) int {
	min := arr[0]
	for _, el := range arr {
		if el < min {
			min = el
		}
	}

	return min
}
func (ListChallenge) FindMinimum2(arr []int) int {
	min := arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
	}

	return min
}

// FindSecondMaximum returns the second largest element in the list.
// input [1, 2, 3, 4]
// output 3
func (ListChallenge) FindSecondMaximum(arr []int) int {
	max := arr[0]
	secondMax := arr[0]

	for _, el := range arr {
		if el > max {
			secondMax = max
			max = el
		} else if el > secondMax {
			secondMax = el
		}
	}

	return secondMax
}
func (ListChallenge) FindSecondMaximum2(arr []int) int {
	max := arr[0]
	secondMax := arr[0]

	for _, val := range arr {
		if val > max {
			secondMax = max
			max = val
		} else if val > secondMax {
			secondMax = val
		}
	}
	return secondMax
}

// RotateToRightK rotate to right a list by k elements
/*
Sample Input [1,2,3,4] k=1
output = [4,1,2,3]
*/
func (ListChallenge) RotateToRightK(arr []int, k int) []int {
	for i := 0; i < k; i++ {
		arr = append(arr[len(arr)-1:], arr[:len(arr)-1]...)
	}

	return arr
}
func (ListChallenge) RotateToRightK2(arr []int, k int) [5]int {
	var arr2 [5]int

	for i := range arr {
		pos := (i + k) % len(arr)
		arr2[pos] = arr[i]
	}

	return arr2
}

// RotateToLeftK rotate to left a list by k elements
/*
Sample Input [1,2,3,4] k=1
output = [2,3,4,1]
*/
func (ListChallenge) RotateToLeftK(arr []int, k int) []int {
	for i := 0; i < k; i++ {
		arr = append(arr[1:], arr[0])
	}

	return arr
}
func (ListChallenge) RotateToLeftK2(arr []int, k int) [5]int {
	var arr2 [5]int
	k = k % len(arr)

	for i := range arr {
		pos := (i - k + len(arr)) % len(arr)
		arr2[pos] = arr[i]
	}
	return arr2
}

// RearrangeArray takes an array and put the negative elements in the back and positive elements in the front
/*
Sample Input [1,-2,3,4]
output = [-2,4,1,3]
*/
func (ListChallenge) RearrangeArray(arr []int) []int {
	res := make([]int, 0)
	for _, el := range arr {
		if el >= 0 {
			res = append(res, el)
		} else {
			res = append([]int{el}, res...)
		}
	}

	return res
}
func (ListChallenge) RearrangeArray2(arr []int) []int {
	var res []int

	for _, val := range arr {
		if val > 0 {
			res = append(res, val)
		} else {
			res = append([]int{val}, res...)
		}
	}

	return res
}

// ArrangeMaxMinElements arrange elements in such a way that the maximum element appears at first position,
//then minimum at second, then second maximum at third and second minimum at fourth and so on.
/*
Sample Input [1,2,3,4,5]
output = [5,1,4,2,3]
*/
func (ListChallenge) ArrangeMaxMinElements(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	res := make([]int, 0)
	minIdx := 0
	maxIdx := len(arr) - 1
	for len(res) != len(arr) {
		res = append(res, arr[maxIdx])
		if len(res) == len(arr) {
			break
		}
		maxIdx--

		res = append(res, arr[minIdx])
		minIdx++
	}

	return res
}
func (ListChallenge) ArrangeMaxMinElements2(arr []int) []int {
	var res []int

	for i := 0; i < len(arr); i++ {
		res = append(res, arr[len(arr)-i-1])

		if len(res) == len(arr) {
			break
		}

		res = append(res, arr[i])
	}

	return res
}

// FindMaxSumSublist finds the maximum sum of a sublist in a given list.
/*
Sample Input [-2,10,7,-5, 15,6]
output =  [10 7 -5 15 6] 33
*/
func (ListChallenge) FindMaxSumSublist(arr []int) ([]int, int) {
	maxSum := arr[0]
	currSum := arr[0]

	sublist := make([]int, 0)
	maxSublist := make([]int, 0)
	for _, el := range arr {
		if currSum+el > el {
			currSum += el
			sublist = append(sublist, el)
		} else {
			currSum = el
			sublist = []int{el}
		}

		if currSum > maxSum {
			maxSum = currSum
			maxSublist = sublist
		}
	}

	return maxSublist, maxSum
}

func (ListChallenge) FindMaxSumSublist2(arr []int) ([]int, int) {
	var res []int
	var maxGeneralSum int
	var maxCurrentSum int

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i, _ := range arr {
		maxCurrentSum += arr[i]
		if maxCurrentSum > arr[i] {
			res = append(res, arr[i])
		} else {
			maxCurrentSum = arr[i]
			res = []int{arr[i]}
		}

		maxGeneralSum = max(maxGeneralSum, maxCurrentSum)
	}

	return res, maxGeneralSum
}
