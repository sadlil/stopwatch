package stopwatch

import (
	"errors"
	"fmt"
	"github.com/sadlil/stopwatch/timeunit"
	"time"
)

type stopWatch struct {
	isRunning   bool
	startTime   time.Time
	elapsedTime time.Duration
	l           laps
}

func newStopWatch(t time.Time, running bool, elapsed int64) *stopWatch {
	return &stopWatch{
		isRunning:   running,
		startTime:   t,
		elapsedTime: time.Duration(elapsed),
		l:           make(laps),
	}
}

func New() *stopWatch {
	return newStopWatch(timeunit.ZeroTime(), false, 0)
}

func NewFrom(t time.Time) *stopWatch {
	e := time.Now().Sub(t).Nanoseconds()
	return newStopWatch(t, false, e)
}

func NewStarted() *stopWatch {
	return newStopWatch(time.Now(), true, 0)
}

func NewStartedFrom(t time.Time) *stopWatch {
	e := time.Now().Sub(t).Nanoseconds()
	return newStopWatch(t, true, e)
}

func (s *stopWatch) IsRunning() bool {
	return s.isRunning
}

func (s *stopWatch) Start() (*stopWatch, error) {
	if s.isRunning {
		return s, errors.New("stopwatch is already running")
	}
	s.isRunning = true
	if s.startTime.IsZero() {
		s.startTime = time.Now()
	}
	return s, nil
}

func (s *stopWatch) Stop() (*stopWatch, error) {
	if !s.isRunning {
		return s, errors.New("stopwatch is already stoped")
	}
	s.isRunning = false
	s.elapsedTime = time.Since(s.startTime)
	return s, nil
}

func (s *stopWatch) Reset() *stopWatch {
	s.isRunning = false
	s.startTime = timeunit.ZeroTime()
	s.elapsedTime = time.Duration(0)
	return s
}

func (s *stopWatch) Elapsed(u timeunit.TimeUnit) int64 {
	return s.elapsed(u)
}

func (s *stopWatch) ElapsedNanos() int64 {
	return s.elapsed(timeunit.NANOSECONDS)
}

func (s *stopWatch) ElapsedTime() time.Duration {
	return s.elapsedTime
}

func (s *stopWatch) ElapsedString() string {
	var duration time.Duration

	if s.isRunning {
		if s.startTime.IsZero() {
			return "0"
		}
		duration = time.Since(s.startTime)
	} else {
		duration = s.elapsedTime
	}
	return duration.String()
}

func (s *stopWatch) ElapsedDuration() time.Duration {
	var duration time.Duration
	if s.isRunning {
		if s.startTime.IsZero() {
			return time.Duration(0)
		}
		duration = time.Since(s.startTime)
	} else {
		duration = s.elapsedTime
	}
	return duration
}

func (s *stopWatch) elapsed(u timeunit.TimeUnit) int64 {
	duration := s.ElapsedDuration()
	return toTimeUnit(duration, u)
}

func (s *stopWatch) PrintString() {
	fmt.Println("Elasped Time:", s.ElapsedString())
}

func (s *stopWatch) Print(u timeunit.TimeUnit) {
	time := s.elapsed(u)
	fmt.Println("Elasped Time:", time, u.Name())
}

func toTimeUnit(duration time.Duration, u timeunit.TimeUnit) int64 {
	nanos := duration.Nanoseconds()
	if u == timeunit.NANOSECONDS {
		return int64(nanos)
	} else if u == timeunit.MICROSECONDS {
		return int64(nanos / 1000)
	} else if u == timeunit.MILLISECONDS {
		return int64(nanos / 1000000)
	} else if u == timeunit.SECONDS {
		return int64(duration.Seconds())
	} else if u == timeunit.MINUTES {
		return int64(duration.Minutes())
	} else if u == timeunit.HOURS {
		return int64(duration.Hours())
	} else if u == timeunit.DAYS {
		return int64(duration.Hours() / 24)
	}
	return int64(0)
}
