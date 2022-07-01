package sortingandsearching

import (
	"sort"
	"strings"
)

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
/*
sample input:[7, 8, 9, 10, 11, 12, 13, 14, 0, 3, 5, 6]
output:8
*/
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
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return 0
}

// BinarySearch search number in a sorted array using binary search
// Time complexity: O(log(n))
// Space complexity: O(1)
/*
sample input:[1, 2, 3, 4, 5] 4
output:3
 */
func (sasc) BinarySearch(arr []int, num int) int {
	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 {
		if arr[0] == num {
			return 0
		}
		return -1
	}

	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == num {
			return mid
		}

		if arr[mid] < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// GroupAnagrams given an array of strings returns groups of anagrams
// Time complexity: O(n)
// Space complexity: O(n)
/*
sample input:["abc", "bca", "cab", "cba", "xyz", "yzx"]
output:[[abc, bca, cab, cba], [xyz, yzx]]
 */
func (s sasc) GroupAnagrams(arr []string) [][]string {
	if len(arr) == 0 {
		return nil
	}

	memo := make(map[string][]string)

	for _, el := range arr {
		key := s.sortString(el)
		memo[key] = append(memo[key], el)
	}

	res := make([][]string, 0)

	for _, el := range memo {
		res = append(res, el)
	}

	return res
}

func (sasc) sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}