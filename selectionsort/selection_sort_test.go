package selectionsort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSelectionSort_Sort(t *testing.T) {
	arr := []int{3, 2, 1}
	expected := []int{1, 2, 3}

	arr = NewSelectionSort().Sort(arr)

	require.Equal(t, expected, arr)
}
