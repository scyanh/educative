package mergeintervals

import "sort"

// MergeAllOverlappingIntervals
/*
Intervals: [[6,7], [2,4], [5,9]]
Output: [[2,4], [5,9]]
Explanation: Since the intervals [6,7] and [5,9] overlap, we merged them into one [5,9].

Intervals: [[1,4], [2,6], [3,5]]
Output: [[1,6]]
Explanation: Since all the given intervals overlap, we merged them into one.
 */
func (mergeIntervals) MergeAllOverlappingIntervals(intervals [][]int) [][]int {
	var res [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start := intervals[0][0]
	end := intervals[0][len(intervals[0])-1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			end= max(intervals[i][len(intervals[i])-1], end)
		} else {
			// push current
			res = append(res, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][len(intervals[i])-1]
		}
	}

	res = append(res, []int{start, end})

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
