package datastructures

import (
	"errors"
)

type listChallenge struct {
}

func NewListChallenge() listChallenge {
	return listChallenge{}
}

func (listChallenge) RemoveEven(arr []int) []int {
	res := make([]int, 0)
	for _, el := range arr {
		if el%2 != 0 {
			res = append(res, el)
		}
	}

	return res
}

func (listChallenge) MergeTwoSortedList(arr1, arr2 []int) []int {
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

// AddTwoNumbersAddToK In this problem, you have to implement the find_sum(lst,k) function
// which will take a number k as input and return two numbers that add up to k.
func (listChallenge) AddTwoNumbersAddToK(arr []int, k int) (int, int, error) {
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
func (listChallenge) FindProductAllElements(arr []int) []int {
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

//FindMinimum finds the smallest number in the given list.
func (listChallenge) FindMinimum(arr []int) int {
	min := arr[0]
	for _, el := range arr {
		if el < min {
			min = el
		}
	}

	return min
}

// FindSecondMaximum returns the second largest element in the list.
func (listChallenge) FindSecondMaximum(arr []int) int {
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

// RotateToRightK rotate to right a list by k elements
/*
Sample Input [1,2,3,4] k=1
output = [4,1,2,3]
*/
func (listChallenge) RotateToRightK(arr []int, k int) []int {
	for i := 0; i < k; i++ {
		arr = append(arr[len(arr)-1:], arr[:len(arr)-1]...)
	}

	return arr
}

// RotateToLeftK rotate to left a list by k elements
/*
Sample Input [1,2,3,4] k=1
output = [2,3,4,1]
*/
func (listChallenge) RotateToLeftK(arr []int, k int) []int {
	for i := 0; i < k; i++ {
		arr = append(arr[1:], arr[0])
	}

	return arr
}

// RearrangeArray takes an array and put the negative elements in the back and positive elements in the front
/*
Sample Input [1,-2,3,4]
output = [-2,4,1,3]
*/
func (listChallenge) RearrangeArray(arr []int) []int {
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

// ArrangeMaxMinElements arrange elements in such a way that the maximum element appears at first position,
//then minimum at second, then second maximum at third and second minimum at fourth and so on.
/*
Sample Input [1,2,3,4,5]
output = [5,1,4,2,3]
*/
func (listChallenge) ArrangeMaxMinElements(arr []int) []int {
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
