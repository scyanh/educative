package mergeintervals

// InsertInterval Given a list of non-overlapping intervals sorted by their start time,
// insert a given interval at the correct position and merge all necessary intervals to produce
// a list that has only mutually exclusive intervals.
/*
Input: Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,6]
Output: [[1,3], [4,7], [8,12]]
Explanation: After insertion, since [4,6] overlaps with [5,7], we merged them into one [4,7].
*/
func (mergeIntervals) InsertInterval(intervals [][]int, newInterval []int) [][]int {
	var res [][]int

	i := 0
	for i < len(intervals) {
		if intervals[i][1] >= newInterval[0] {
			mergeInterval(intervals[i], newInterval)
			break
		}

		res = append(res, intervals[i])
		i++
	}

	for ; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}

	return res
}

func mergeInterval(current []int, newInterval []int) {
	current[0] = min(current[0], newInterval[0])
	current[1] = max(current[1], newInterval[1])
}
