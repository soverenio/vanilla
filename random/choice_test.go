package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapChoice(t *testing.T) {
	a := map[string]int{}

	for i := 0; i < 1000; i++ {
		_, _, ok := MapChoice(a)
		assert.False(t, ok)
	}

	a["a"] = 1

	for i := 0; i < 1000; i++ {
		k, v, ok := MapChoice(a)
		assert.True(t, ok)
		assert.Equal(t, "a", k)
		assert.Equal(t, 1, v)
	}

	a["b"] = 2

	for i := 0; i < 1000; i++ {
		k, v, ok := MapChoice(a)
		assert.True(t, ok)
		assert.Contains(t, []string{"a", "b"}, k)
		assert.Contains(t, []int{1, 2}, v)
	}
}
