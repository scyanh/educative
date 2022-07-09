package twopointers

// PairWithTargetSum returns the first pair of integers in the array that sum to the target sum
// Time Complexity: O(n)
// Space Complexity: O(1)
// input: arr = [2, 5, 9, 11]
// target = 11
// output: [0, 2]
// Explanation: The numbers at index 0 and 2 add up to 11: 2+9=11
func (twoPointers) PairWithTargetSum(arr []int, target int) []int {
	left := 0
	right := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		sum := arr[left] + arr[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left, right}
		}
	}

	return []int{}
}
