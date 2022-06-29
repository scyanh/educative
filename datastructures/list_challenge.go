package datastructures

import "errors"

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
