package slidingwindow

// MaxSubarrayWithMaxK max subarray with maximum k elements or less
// input: arr = [8, -7, 8, 3], k = 3
// output: 11
// complexity: O(n^2)
func (s slidingWindowChallenge) MaxSubarrayWithMaxK(arr []int, k int) int {
	res := arr[0]
	var left int
	maxEndingHere := arr[0]

	for i := 1; i < len(arr); i++ {
		if i-left+1 == k {
			maxEndingHere = s.getMaxFromSlice(arr[left:i+1])
			left++
		}else{
			maxEndingHere = s.getMaxFromSlice(arr[left:i+1])
		}

		res = max(maxEndingHere, res)
	}

	return res
}

func (slidingWindowChallenge) getMaxFromSlice(arr []int) int {
	res := arr[0]
	maxEndingHere := arr[0]

	for i := 1; i < len(arr); i++ {
		maxEndingHere = max(maxEndingHere+arr[i], arr[i])
		res = max(maxEndingHere, res)
	}

	return res
}
