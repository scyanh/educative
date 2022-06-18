package insertionsort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsertionSort_Sort(t *testing.T) {
	arr := []int{3, 2, 1}
	expected := []int{1, 2, 3}

	arr = NewInsertionSort().Sort(arr)

	require.Equal(t, expected, arr)
}
