package stopwatch

import (
	"fmt"
	"github.com/sadlil/stopwatch/timeunit"
	"testing"
	"time"
)

func TestStopWatch(t *testing.T) {
	first := New()
	first.Start()
	fmt.Println(first.ElapsedNanos())
	fmt.Println(first.ElapsedString())
	time.Sleep(time.Second)
	fmt.Println(first.ElapsedNanos())
	fmt.Println(first.Elapsed(timeunit.DAYS),
		first.Elapsed(timeunit.HOURS),
		first.Elapsed(timeunit.MINUTES),
		first.Elapsed(timeunit.SECONDS),
		first.Elapsed(timeunit.MILLISECONDS),
		first.Elapsed(timeunit.MICROSECONDS),
		first.Elapsed(timeunit.NANOSECONDS),
	)
	fmt.Println(first.ElapsedString())
}
