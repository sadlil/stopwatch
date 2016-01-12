package timeunit

type TimeUnit int

const (
	DAYS TimeUnit = iota
	HOURS
	MICROSECONDS
	MILLISECONDS
	MINUTES
	NANOSECONDS
	SECONDS
)
