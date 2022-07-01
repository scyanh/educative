package sortingandsearching

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSasc_FindPivotIndex(t *testing.T) {
	arr := []int{7, 8, 9, 10, 11, 12, 13, 14, 0, 3, 5, 6}

	pivot := NewSASC().FindPivotIndex(arr)
	expected:=8

	require.Equal(t, expected, pivot)
}
