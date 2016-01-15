package stopwatch

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestLaps(t *testing.T) {
	s := New()
	err := s.AddLap("hello")
	assert.NotEqual(t, err, nil)

	s.Start()
	s.AddLap("lap-1")
	s.AddLap("lap-2")

	ti, err := s.GetLapNanos("lap-1")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, ti, 0)

	ti, err = s.GetLapNanos("lapnotfound")
	assert.NotEqual(t, err, nil)
	assert.Equal(t, ti, int64(0))

	s.PrintAllLapString()
}
