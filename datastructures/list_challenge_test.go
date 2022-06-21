package datastructures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListChallenge_RemoveEven(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	expected := []int{1, 3, 5}

	res := NewListChallenge().RemoveEven(arr)

	require.Equal(t, expected, res)
}

func TestListChallenge_MergeTwoSortedList(t *testing.T) {
	arr := []int{1, 2, 5, 6}
	arr2 := []int{3, 4}
	expected := []int{1, 2, 3, 4, 5, 6}

	res := NewListChallenge().MergeTwoSortedList(arr, arr2)
	require.Equal(t, expected, res)
}
