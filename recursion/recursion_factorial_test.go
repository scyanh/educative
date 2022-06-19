package recursion

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRecursiveFactorial_Factorial(t *testing.T) {
	res := NewRecursiveFactorial().Factorial(5)
	expected:=120
	require.Equal(t, expected, res)
}
