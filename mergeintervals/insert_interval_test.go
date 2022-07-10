package mergeintervals

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeIntervals_InsertInterval(t *testing.T) {
	arr := [][]int{
		{1, 3},
		{5, 7},
		{8, 12},
	}

	newInterval := []int{4, 6}
	expected := [][]int{
		{1, 3},
		{4, 7},
		{8, 12},
	}
	res := NewMergeIntervals().InsertInterval(arr, newInterval)

	require.Equal(t, expected, res)
}
