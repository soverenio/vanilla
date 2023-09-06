package zero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsZero(t *testing.T) {
	var a [16]byte
	assert.True(t, IsZero(a))

	a[0] = 100
	assert.False(t, IsZero(a))
}
