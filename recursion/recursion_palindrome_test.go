package recursion

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRecursionPalindrome_IsPalindrome(t *testing.T) {
	r := NewRecursionPalindrome()
	res1 := r.IsPalindrome("abba")
	res2 := r.IsPalindrome("abbac")

	require.Equal(t, true, res1)
	require.Equal(t, false, res2)
}
