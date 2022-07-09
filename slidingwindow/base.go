package slidingwindow

type slidingWindowChallenge struct {
}

func NewSlidingWindowChallenge() *slidingWindowChallenge {
	return &slidingWindowChallenge{}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}