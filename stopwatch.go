package stopwatch

import (
	"errors"
	"github.com/sadlil/stopwatch/timeunit"
	"time"
)

type stopWatch struct {
	isRunning   bool
	startTime   time.Time
	elapsedTime time.Duration
	laps        *map[string]time.Time
}

func newStopWatch(t time.Time, running bool) *stopWatch {
	return &stopWatch{
		isRunning:   running,
		startTime:   t,
		elapsedTime: time.Duration(0),
		laps:        new(map[string]time.Time),
	}
}

func New() *stopWatch {
	return newStopWatch(timeunit.ZeroTime(), false)
}

func NewFrom(t time.Time) *stopWatch {
	return newStopWatch(t, false)
}

func NewStarted() *stopWatch {
	return newStopWatch(timeunit.ZeroTime(), true)
}

func NewStartedFrom(t time.Time) *stopWatch {
	return newStopWatch(t, true)
}

func (s *stopWatch) IsRunning() bool {
	return s.isRunning
}

func (s *stopWatch) Start() (*stopWatch, error) {
	if s.isRunning {
		return s, errors.New("stopwatch is already running")
	}
	s.isRunning = true
	s.startTime = time.Now()
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

func (s *stopWatch) Elapsed(unit timeunit.TimeUnit) int64 {
	return s.elapsed(unit)
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

func (s *stopWatch) elapsed(u timeunit.TimeUnit) int64 {
	var duration time.Duration

	if s.isRunning {
		if s.startTime.IsZero() {
			return 0
		}
		duration = time.Since(s.startTime)
	} else {
		duration = s.elapsedTime
	}
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
	return 0
}
