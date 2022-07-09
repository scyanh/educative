package slidingwindow

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSlidingWindowChallenge(t *testing.T) {
	slidingWindowChallenge := NewSlidingWindowChallenge()
	require.NotNil(t, slidingWindowChallenge)

	arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5
	res := slidingWindowChallenge.FindAverageOfSubarrayOfSizeK(arr, k)
	require.Equal(t, []float64{2.2, 2.8, 2.4, 3.6, 2.8}, res)
}
