package slidingwindow

// MaxSubarrayOfSizeK max subarray of k
// Given an array of integers and a number k, find the maximum subarray of k consecutive elements in the array.
// Hint. Subtract the element going out of the sliding window, i.e., subtract the first element of the window.
// Add the new element getting included in the sliding window, i.e., the element coming right after the end of the window.
//
// Example 1:
// Input: arr = [-3,4,3,-2,2,5], k = 5
// Output: 8
// Example 2:
// Input: arr = [2, 1, 5, 1, 3, 2], k = 3
// Output: 9
func (slidingWindowChallenge) MaxSubarrayOfSizeK(arr []int, k int) int {
	var res int
	var sum int
	var start int

	for i:=0; i<len(arr); i++ {
		sum += arr[i]

		if i-start+1 == k {
			res = max(res, sum)
			sum -= arr[start]
			start++
		}
	}

	return res
}
