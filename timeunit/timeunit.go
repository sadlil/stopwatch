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

func ZeroTime() time.Time {
	zTime := time.Time{}
	return zTime
}
