// Package stopwatch provides a timer that implements common stopwatch
// functionality.
package stopwatch

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	start, stop, lap time.Time
	laps             []time.Duration
}

// New creates a new Stopwatch. To start the stopwatch Start() should be invoked.
func New() *Stopwatch {
	return &Stopwatch{}
}

// Starts creates a new Stopwatch that starts the timer immediately.
func Start() *Stopwatch {
	s := &Stopwatch{}
	s.init()
	return s
}

func (s *Stopwatch) init() {
	s.start, s.lap = time.Now(), time.Now()
	s.laps = make([]time.Duration, 0)
}

func (s *Stopwatch) isStopped() bool {
	return s.stop.After(s.start)
}

func (s *Stopwatch) isResetted() bool {
	return s.start.IsZero()
}

// ElapsedTime returns the duration between the start and current time.
func (s *Stopwatch) ElapsedTime() time.Duration {
	if s.isStopped() {
		return s.stop.Sub(s.start)
	}

	if s.isResetted() {
		return time.Duration(0)
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
	if s.isResetted() {
		s.init()
	} else { //stopped
		s.start = s.start.Add(time.Since(s.stop))
	}
}

// Reset resets the timer. It needs to be started again with the Start() method
func (s *Stopwatch) Reset() {
	s.start, s.stop, s.lap = time.Time{}, time.Time{}, time.Time{}
	s.laps = nil
}

// Lap takes and stores the current lap time and returns the elapsed time
// since the latest lap.
func (s *Stopwatch) Lap() time.Duration {
	// There is no lap if the timer is resetted or stoped
	if s.isStopped() || s.isResetted() {
		return time.Duration(0)
	}

	lap := time.Since(s.lap)
	s.lap = time.Now()
	s.laps = append(s.laps, lap)

	return lap
}

// Laps returns the list of all laps
func (s *Stopwatch) Laps() []time.Duration {
	return s.laps
}

// String representation of a single Stopwatch instance.
func (s *Stopwatch) String() string {
	return fmt.Sprintf("[start: %s current: %s elapsed: %s]\n",
		s.start.Format(time.Stamp), time.Now().Format(time.Stamp), s.ElapsedTime())
}
