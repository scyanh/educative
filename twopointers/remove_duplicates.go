package twopointers

// RemoveDuplicates given an array of sorted numbers, removes all duplicate number instances
// from it in-place, such that each element appears only once. The relative order of the elements
// should be kept the same and you should not use any extra space so that that the solution have
// a space complexity of O(1).
//
// Move all the unique elements at the beginning of the array and after moving return
// the length of the subarray that has no duplicate in it.
func (twoPointers) RemoveDuplicates(arr []int) int {
	nextNonDuplicate := 1
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr[nextNonDuplicate-1] {
			arr[nextNonDuplicate] = arr[i]
			nextNonDuplicate++
		}
	}
	return nextNonDuplicate
}
