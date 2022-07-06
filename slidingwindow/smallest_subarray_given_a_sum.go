package slidingwindow

import "math"

// FindSmallestSubarrayWithGivenSum finds the smallest subarray with a greater sum
/*
Example 1:
Input: arr = [2,3,1,2,4,3], k = 7
Output: 2
*/
func (slidingWindowChallenge) FindSmallestSubarrayWithGivenSum(nums []int, s int) int {
	n := len(nums) // length of array
	if n == 0 {
		return 0 // if array is empty, return 0
	}

	left := 0 // left pointer
	res := math.MaxInt // result
	sum := 0 // sum of subarray

	for i := 0; i < n; i++ { // iterate through array
		sum += nums[i] // add new element to sum
		for sum >= s { // if sum is greater than or equal to s
			res = min(res, i + 1 - left) // update result
			sum -= nums[left] 		   // remove the first element
			left++ 					   // move the left pointer
		}
	}

	if res != math.MaxInt{ // if result is not max int
		return res // return result
	}

	return 0 // if result is max int, return 0
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

