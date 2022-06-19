package recursion

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRecursionPower_Power(t *testing.T) {
	res1 := NewRecursionPower().Power(2, 3)
	res2 := NewRecursionPower().Power(2, 6)

	require.Equal(t, 8, res1)
	require.Equal(t, 64, res2)
}
