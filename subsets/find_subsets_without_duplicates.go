package subsets

import "sort"

func (subset) FindSubsetsWithoutDuplicates(arr []int) [][]int {
	subsets := [][]int{{}}

	sort.Ints(arr)

	for j, el := range arr {
		n := len(subsets)
		for i := 0; i < n; i++ {
			set := subsets[i]
			if j>0 && el==arr[j-1]{
				if i==0 || i==1{
					continue
				}
			}
			set = append(set, el)
			subsets = append(subsets, set)
		}

	}

	return subsets
}
