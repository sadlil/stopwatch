package stopwatch

import (
	"errors"
	"fmt"
	"github.com/sadlil/stopwatch/timeunit"
	"time"
)

type laps map[string]time.Duration

type lapsDesc struct {
	Name     string
	Duration time.Duration
}

func (s *stopWatch) AddLap(name string) error {
	if !s.isRunning {
		return errors.New("stopwatch not running")
	}
	s.l[name] = s.ElapsedDuration()
	return nil
}

func (s *stopWatch) GetLap(name string, u timeunit.TimeUnit) (int64, error) {
	d, ok := s.l[name]
	if !ok {
		return 0, errors.New("no lap found with name")
	}
	return toTimeUnit(d, u), nil
}

func (s *stopWatch) GetLapString(name string) (string, error) {
	d, ok := s.l[name]
	if !ok {
		return "0", errors.New("no lap found with name")
	}
	return d.String(), nil
}

func (s *stopWatch) GetLapNanos(name string) (int64, error) {
	d, ok := s.l[name]
	if !ok {
		return 0, errors.New("no lap found with name")
	}
	return d.Nanoseconds(), nil
}

func (s *stopWatch) GetLaps() []lapsDesc {
	lapsList := make([]lapsDesc, len(s.l))
	i := 0
	for k, v := range s.l {
		lapsList[i].Name = k
		lapsList[i].Duration = v
		i++
	}
	return lapsList
}

func (s *stopWatch) PrintLap(name string, u timeunit.TimeUnit) {
	fmt.Println("Lap", name, "time", toTimeUnit(s.l[name], u), u.Name())
}

func (s *stopWatch) PrintLapString(name string) {
	fmt.Println("Lap", name, "time", s.l[name].String())
}

func (s *stopWatch) PrintAllLapString() {
	for k, v := range s.l {
		fmt.Println("Lap", k, "time", v.String())
	}
}

func (s *stopWatch) PrintAllLap(u timeunit.TimeUnit) {
	for k, v := range s.l {
		fmt.Println("Lap", k, "time", toTimeUnit(v, u), u.Name())
	}
}
