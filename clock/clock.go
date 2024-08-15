package clock

import "time"

type Clock interface {
	Now() time.Time
}

var _ Clock = TimeClock{}

type TimeClock struct{}

func NewTimeClock() TimeClock {
	return TimeClock{}
}

func (TimeClock) Now() time.Time {
	return time.Now()
}
