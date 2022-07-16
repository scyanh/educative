package subsets

func (subset) FindSubsets(arr []int) [][]int {
	var subsets [][]int

	subsets = append(subsets, []int{})

	for _, el := range arr {
		n := len(subsets)

		for i := 0; i < n; i++ {
			set := subsets[i]
			set = append(set, el)
			subsets = append(subsets, set)
		}
	}

	return subsets
}
