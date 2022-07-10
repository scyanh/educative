package mergeintervals

type mergeIntervals struct{

}

func NewMergeIntervals() mergeIntervals{
	return mergeIntervals{}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}