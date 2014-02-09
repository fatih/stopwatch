// Package stopwatch provides a timer that implements common stopwatch
// functionality.
package stopwatch

import (
	"fmt"
	"time"
)

type StopWatch struct {
	start time.Time
}

// New creates a new StopWatch that starts the timer immediately.
func New() *StopWatch {
	return &StopWatch{
		start: time.Now(),
	}
}

// ElapsedTime returns the duration between the start and current time.
func (s *StopWatch) ElapsedTime() time.Duration {
	return time.Since(s.start)
}

func (s *StopWatch) Stop()  {}
func (s *StopWatch) Pause() {}

// Reset resets the timer. It needs to be started again with the Start() method
func (s *StopWatch) Reset() {
	s.start = time.Time{}
}

func (s *StopWatch) Lap() {}

func (s *StopWatch) String() string {
	return fmt.Sprintf("[start: %s current: %s elapsed: %s]\n",
		s.start.Format(time.Stamp), time.Now().Format(time.Stamp), s.ElapsedTime())
}
