package timeunit

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestZeroTime(t *testing.T) {
	z := ZeroTime()
	is := z.IsZero()
	assert.Equal(t, z, ZeroTime())
	assert.Equal(t, true, is)
}
