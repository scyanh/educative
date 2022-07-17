package subsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubset_GeneratePermutations(t *testing.T) {
	arr := []int{1, 3, 5}

	res := NewSubset().GeneratePermutations(arr)
	expected := [][]int{
		{5, 3, 1}, {3, 5, 1}, {3, 1, 5}, {5, 1, 3}, {1, 5, 3}, {1, 3, 5},
	}

	assert.ElementsMatch(t, expected, res)
}
