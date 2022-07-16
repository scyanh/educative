package subsets

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubset_FindSubsets(t *testing.T) {
	arr := []int{1, 2, 3}
	res:= NewSubset().FindSubsets(arr)

	expected:= [][]int{{},{1}, {2}, {1, 2}, {3}, {1, 3}, {2,3}, {1,2,3}}

	require.Equal(t, expected, res)
}
