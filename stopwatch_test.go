package stopwatch

import (
	"fmt"
	"github.com/sadlil/stopwatch/timeunit"
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"time"
)

func TestStopWatch(t *testing.T) {
	fmt.Println("Testing StopWatch New")
	first := New()
	assert.Equal(t, first.IsRunning(), false)
	assert.Equal(t, first.startTime, time.Time{})
	assert.Equal(t, first.ElapsedString(), "0")
	assert.Equal(t, first.ElapsedNanos(), int64(0))
	assert.Equal(t, first.ElapsedTime(), time.Duration(0))

	a, err := first.Start()
	assert.Equal(t, err, nil)
	assert.Equal(t, a.IsRunning(), true)
	assert.Equal(t, first.IsRunning(), true)

	a, err = first.Start()
	assert.NotEqual(t, err, nil)
	assert.Equal(t, a.IsRunning(), true)
	assert.Equal(t, first.IsRunning(), true)

	first.PrintString()
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.MICROSECONDS)
	first.Print(timeunit.MILLISECONDS)
	first.Print(timeunit.SECONDS)

	a, err = first.Stop()
	assert.Equal(t, err, nil)
	assert.Equal(t, a.IsRunning(), false)
	assert.Equal(t, first.IsRunning(), false)

	assert.Equal(t, first.ElapsedNanos(), first.ElapsedNanos())
	assert.Equal(t, first.ElapsedString(), first.ElapsedString())

	first.PrintString()
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.MICROSECONDS)
	first.Print(timeunit.MILLISECONDS)
	first.Print(timeunit.SECONDS)

	a, err = first.Stop()
	assert.NotEqual(t, err, nil)
	assert.Equal(t, a.IsRunning(), false)
	assert.Equal(t, first.IsRunning(), false)

	assert.Equal(t, first.ElapsedNanos(), first.ElapsedNanos())
	assert.Equal(t, first.ElapsedString(), first.ElapsedString())

	a = first.Reset()
	assert.Equal(t, a, first)
	assert.Equal(t, a.IsRunning(), false)
	assert.Equal(t, a.ElapsedString(), "0")
	assert.Equal(t, a.startTime, time.Time{})

	first.PrintString()
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.MICROSECONDS)
	first.Print(timeunit.MILLISECONDS)
	first.Print(timeunit.SECONDS)
}

func TestStopWatchNewStarted(t *testing.T) {
	fmt.Println("Testing StopWatch NewStarted")
	first := NewStarted()
	assert.Equal(t, first.IsRunning(), true)
	assert.NotEqual(t, first.ElapsedString(), "0")
	assert.NotEqual(t, first.ElapsedNanos(), int64(0))

	a, err := first.Start()
	assert.NotEqual(t, err, nil)
	assert.Equal(t, a.IsRunning(), true)
	assert.Equal(t, first.IsRunning(), true)

	first.PrintString()
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.MICROSECONDS)
	first.Print(timeunit.MILLISECONDS)
	first.Print(timeunit.SECONDS)

	a, err = first.Stop()
	assert.Equal(t, err, nil)
	assert.Equal(t, a.IsRunning(), false)
	assert.Equal(t, first.IsRunning(), false)

	assert.Equal(t, first.ElapsedNanos(), first.ElapsedNanos())
	assert.Equal(t, first.ElapsedString(), first.ElapsedString())

	first.PrintString()
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.MICROSECONDS)
	first.Print(timeunit.MILLISECONDS)
	first.Print(timeunit.SECONDS)

	a, err = first.Stop()
	assert.NotEqual(t, err, nil)
	assert.Equal(t, a.IsRunning(), false)
	assert.Equal(t, first.IsRunning(), false)

	assert.Equal(t, first.ElapsedNanos(), first.ElapsedNanos())
	assert.Equal(t, first.ElapsedString(), first.ElapsedString())

	a = first.Reset()
	assert.Equal(t, a, first)
	assert.Equal(t, a.IsRunning(), false)
	assert.Equal(t, a.ElapsedString(), "0")
	assert.Equal(t, a.startTime, time.Time{})

	first.PrintString()
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.MICROSECONDS)
	first.Print(timeunit.MILLISECONDS)
	first.Print(timeunit.SECONDS)
}

func TestStopWatchNewFrom(t *testing.T) {
	fmt.Println("Testing StopWatch NewFrom")
	first := NewFrom(time.Now().Add(time.Duration(time.Nanosecond * 500)))
	first.Print(timeunit.NANOSECONDS)
	assert.Equal(t, first.IsRunning(), false)

	a, err := first.Start()
	assert.Equal(t, err, nil)
	assert.Equal(t, a.IsRunning(), true)
	assert.Equal(t, first.IsRunning(), true)

	a, err = first.Start()
	assert.NotEqual(t, err, nil)
	assert.Equal(t, a.IsRunning(), true)
	assert.Equal(t, first.IsRunning(), true)

	first.PrintString()
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.MICROSECONDS)
	first.Print(timeunit.MILLISECONDS)
	first.Print(timeunit.SECONDS)

	a, err = first.Stop()
	assert.Equal(t, err, nil)
	assert.Equal(t, a.IsRunning(), false)
	assert.Equal(t, first.IsRunning(), false)

	assert.Equal(t, first.ElapsedNanos(), first.ElapsedNanos())
	assert.Equal(t, first.ElapsedString(), first.ElapsedString())

	first.PrintString()
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.MICROSECONDS)
	first.Print(timeunit.MILLISECONDS)
	first.Print(timeunit.SECONDS)

	a, err = first.Stop()
	assert.NotEqual(t, err, nil)
	assert.Equal(t, a.IsRunning(), false)
	assert.Equal(t, first.IsRunning(), false)

	assert.Equal(t, first.ElapsedNanos(), first.ElapsedNanos())
	assert.Equal(t, first.ElapsedString(), first.ElapsedString())

	a = first.Reset()
	assert.Equal(t, a, first)
	assert.Equal(t, a.IsRunning(), false)
	assert.Equal(t, a.ElapsedString(), "0")
	assert.Equal(t, a.startTime, time.Time{})

	first.PrintString()
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.MICROSECONDS)
	first.Print(timeunit.MILLISECONDS)
	first.Print(timeunit.SECONDS)
}

func TestStopWatchNewStartedFrom(t *testing.T) {
	fmt.Println("Testing StopWatch NewStartedFrom")
	first := NewStartedFrom(time.Now().Add(time.Duration(5)))
	assert.Equal(t, first.IsRunning(), true)

	a, err := first.Start()
	assert.NotEqual(t, err, nil)
	assert.Equal(t, a.IsRunning(), true)
	assert.Equal(t, first.IsRunning(), true)

	first.PrintString()
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.MICROSECONDS)
	first.Print(timeunit.MILLISECONDS)
	first.Print(timeunit.SECONDS)

	a, err = first.Stop()
	assert.Equal(t, err, nil)
	assert.Equal(t, a.IsRunning(), false)
	assert.Equal(t, first.IsRunning(), false)

	assert.Equal(t, first.ElapsedNanos(), first.ElapsedNanos())
	assert.Equal(t, first.ElapsedString(), first.ElapsedString())

	first.PrintString()
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.MICROSECONDS)
	first.Print(timeunit.MILLISECONDS)
	first.Print(timeunit.SECONDS)

	a, err = first.Stop()
	assert.NotEqual(t, err, nil)
	assert.Equal(t, a.IsRunning(), false)
	assert.Equal(t, first.IsRunning(), false)

	assert.Equal(t, first.ElapsedNanos(), first.ElapsedNanos())
	assert.Equal(t, first.ElapsedString(), first.ElapsedString())

	a = first.Reset()
	assert.Equal(t, a, first)
	assert.Equal(t, a.IsRunning(), false)
	assert.Equal(t, a.ElapsedString(), "0")
	assert.Equal(t, a.startTime, time.Time{})

	first.PrintString()
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.NANOSECONDS)
	first.Print(timeunit.MICROSECONDS)
	first.Print(timeunit.MILLISECONDS)
	first.Print(timeunit.SECONDS)
}
