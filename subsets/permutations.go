package subsets

func (subset) GeneratePermutations(arr []int) [][]int {
	var res [][]int

	res = generatePermutationsRecursive(arr, 0, []int{}, res)

	return res
}

// generatePermutationsRecursive
func generatePermutationsRecursive(arr []int, idx int, currentPermutation []int, res [][]int) [][]int {
	if idx == len(arr) {
		res = append(res, currentPermutation)
		return res
	}

	for i := 0; i <= len(currentPermutation); i++ {
		// copy currentPermutation
		newPermutation := make([]int, len(currentPermutation))
		copy(newPermutation, currentPermutation)
		// add element to newPermutation at every position
		newPermutation = append(newPermutation[:i], append([]int{arr[idx]}, newPermutation[i:]...)...)
		// recurse
		res = generatePermutationsRecursive(arr, idx+1, newPermutation, res)
	}

	return res
}
