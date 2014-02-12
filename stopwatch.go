// Package stopwatch provides a timer that implements common stopwatch
// functionality.
package stopwatch

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Stopwatch implements the stopwatch functionality. It is not threadsafe by
// design and should be protected when there is a need for.
type Stopwatch struct {
	start, stop, lap time.Time
	laps             []time.Duration
}

// New creates a new Stopwatch. To start the stopwatch Start() should be invoked.
func New() *Stopwatch {
	return &Stopwatch{
		laps: make([]time.Duration, 0),
	}
}

// Start creates a new stopwatch with starting time offset by a user defined
// value. Negative offsets result in a countdown prior to the start of the
// stopwatch. A zero offset starts the stopwatch immediately.
func Start(offset time.Duration) *Stopwatch {
	t := time.Now().Add(offset)
	s := &Stopwatch{
		start: t,
		lap:   t,
		laps:  make([]time.Duration, 0),
	}
	return s
}

// IsStopped shows whether the stopwatch is stopped or not.
func (s *Stopwatch) IsStopped() bool { return s.stop.After(s.start) }

// IsReseted shows whether the stopwatch is reseted or not.
func (s *Stopwatch) IsReseted() bool { return s.start.IsZero() }

// ElapsedTime returns the duration between the start and current time.
func (s *Stopwatch) ElapsedTime() time.Duration {
	if s.IsStopped() {
		return s.stop.Sub(s.start)
	}

	if s.IsReseted() {
		return time.Duration(0)
	}

	return time.Since(s.start)
}

// Print calls fmt.Printf() with the given string and the elapsed time attached.
// Useful to use with a defer statement.
// Example : defer Start().Print("myFunction")
// Output  :  myFunction - elapsed: 2.000629842s
func (s *Stopwatch) Print(msg string) {
	fmt.Printf("%s - elapsed: %s\n", msg, s.ElapsedTime())
}

// Log calls log.Printf() with the given string and the elapsed time attached.
// Useful to use with a defer statement.
// Example : defer Start().Log("myFunction")
// Output: 2014/02/10 00:44:56 myFunction - elapsed: 2.000169591s
func (s *Stopwatch) Log(msg string) {
	log.Printf("%s - elapsed: %s\n", msg, s.ElapsedTime())
}

// Stop stops the timer. To resume the timer Start() needs to be called again.
func (s *Stopwatch) Stop() {
	s.stop = time.Now()
}

// Start resumes or starts the timer. If a Stop() was invoked it resumes the
// timer. If a Reset() was invoked it starts a new session with the given
// offset.
func (s *Stopwatch) Start(offset time.Duration) {
	if s.IsReseted() {
		*s = *Start(offset)
	} else { //stopped
		s.start = s.start.Add(time.Since(s.stop))
	}
}

// Reset resets the timer. It needs to be started again with the Start()
// method.
func (s *Stopwatch) Reset() {
	s.start, s.stop, s.lap = time.Time{}, time.Time{}, time.Time{}
	s.laps = nil
}

// Lap takes and stores the current lap time and returns the elapsed time
// since the latest lap.
func (s *Stopwatch) Lap() time.Duration {
	// There is no lap if the timer is resetted or stoped
	if s.IsStopped() || s.IsReseted() {
		return time.Duration(0)
	}

	lap := time.Since(s.lap)
	s.lap = time.Now()
	s.laps = append(s.laps, lap)

	return lap
}

// Laps returns a slice of all completed laps.
func (s *Stopwatch) Laps() []time.Duration {
	laps := make([]time.Duration, len(s.laps))
	copy(laps, s.laps)
	return laps
}

// String representation of a single Stopwatch instance.
func (s *Stopwatch) String() string {
	return fmt.Sprintf("[start: %s current: %s elapsed: %s]",
		s.start.Format(time.Stamp), time.Now().Format(time.Stamp), s.ElapsedTime())
}

// MarshalJSON implements the json.Marshaler interface. The elapsed time is
// quoted as a string and is in the form "72h3m0.5s". For more info please
// refer to time.Duration.String().
func (s *Stopwatch) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.ElapsedTime().String() + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface. The elapsed time
// is expected to be a string that can be successful parsed with
// time.ParseDuration.
func (s *Stopwatch) UnmarshalJSON(data []byte) (err error) {
	unquoted := strings.Replace(string(data), "\"", "", -1)
	d, err := time.ParseDuration(unquoted)
	if err != nil {
		return err
	}

	// set the start time based on the elapsed time
	s.start = time.Now().Add(-d)
	return nil
}
