// Package stopwatch provides a timer that implements common stopwatch
// functionality.
package stopwatch

import "time"

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
func (s *StopWatch) Reset() {}
func (s *StopWatch) Lap()   {}
