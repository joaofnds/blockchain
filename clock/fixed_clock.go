package clock

import "time"

var _ Clock = FixedClock{}

type FixedClock struct {
	now time.Time
}

func NewFixedClock(now time.Time) FixedClock {
	return FixedClock{now}
}

func (c FixedClock) Now() time.Time {
	return c.now
}
