// Package stopwatch implements a timer that implements a
// stopwatch functionality.
package stopwatch

import "time"

type StopWatch struct{}

func New() *StopWatch                           { return &StopWatch{} }
func (s *StopWatch) ElapsedTime() time.Duration { return 0 }
func (s *StopWatch) Stop()                      {}
func (s *StopWatch) Pause()                     {}
func (s *StopWatch) Reset()                     {}
func (s *StopWatch) Lap()                       {}
