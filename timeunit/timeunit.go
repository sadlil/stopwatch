package timeunit

import "time"

type TimeUnit int

const (
	NANOSECONDS TimeUnit = iota
	MICROSECONDS
	MILLISECONDS
	SECONDS
	MINUTES
	HOURS
	DAYS
)

var timeUnitMap map[TimeUnit]string

func init() {
	timeUnitMap = map[TimeUnit]string{
		NANOSECONDS:  "nanoseconds",
		MICROSECONDS: "microseconds",
		MILLISECONDS: "milliseconds",
		SECONDS:      "seconds",
		MINUTES:      "minutes",
		HOURS:        "hours",
		DAYS:         "days",
	}
}

func (t TimeUnit) Name() string {
	return timeUnitMap[t]
}

func ZeroTime() time.Time {
	zTime := time.Time{}
	return zTime
}
