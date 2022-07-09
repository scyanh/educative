package slidingwindow

// FindAverageOfSubarrayOfSizeK finds average of subarray of size k
/*
Example 1:
Input: arr = [1, 3, 2, 6, -1, 4, 1, 8, 2], k = 5
Output: [2.2, 2.8, 2.4, 3.6, 2.8]
*/
func (slidingWindowChallenge) FindAverageOfSubarrayOfSizeK(arr []int, k int) []float64 {
	var res []float64
	var sum int   // sum of subarray of size k
	var start int // start index of subarray of size k
	var end int   // end index of subarray of size k

	for end < len(arr) {
		sum += arr[end] // add new element to sum
		end++           // move end pointer to next element

		if end-start == k { // sliding window is full
			res = append(res, float64(sum)/float64(k)) // append average of subarray of size k
			sum -= arr[start]                          // remove the first element
			start++                                    // move the start pointer
		}
	}

	return res
}
