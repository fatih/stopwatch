// Package stopwatch provides a timer that implements common stopwatch
// functionality.
package stopwatch

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	start time.Time
	stop  time.Time
}

// New creates a new Stopwatch that starts the timer immediately.
func New() *Stopwatch {
	return &Stopwatch{
		start: time.Now(),
	}
}

// ElapsedTime returns the duration between the start and current time.
func (s *Stopwatch) ElapsedTime() time.Duration {
	if s.stop.After(s.start) {
		return s.stop.Sub(s.start)
	}

	return time.Since(s.start)
}

// Stop stops the timer. To resume the timer Start() needs to be called again.
func (s *Stopwatch) Stop() {
	s.stop = time.Now()
}

// Start resumes or starts the timer. If a Stop() was invoked it resumes the
// timer. If a Reset() was invoked it starts a new session.
func (s *Stopwatch) Start() {
	if s.start.IsZero() { // reseted
		s.start = time.Now()
	} else { //stopped
		s.start = s.start.Add(time.Since(s.stop))
	}
}

// Reset resets the timer. It needs to be started again with the Start() method
func (s *Stopwatch) Reset() {
	s.start = time.Time{}
}

func (s *Stopwatch) Lap() {}

func (s *Stopwatch) String() string {
	return fmt.Sprintf("[start: %s current: %s elapsed: %s]\n",
		s.start.Format(time.Stamp), time.Now().Format(time.Stamp), s.ElapsedTime())
}
