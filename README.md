# Stopwatch [![GoDoc](https://godoc.org/github.com/fatih/stopwatch?status.png)](http://godoc.org/github.com/fatih/stopwatch) [![Build Status](https://travis-ci.org/fatih/stopwatch.png)](https://travis-ci.org/fatih/stopwatch)

Stopwatch implements a common set of functions of a stop-watch. It's handy for
calculating process times between function calls, logging times and many other
use cases. Feel free to fork and send a pull request for any
changes/improvements. For usage see examples below or click on the godoc
badge.

## Install

```bash
go get github.com/fatih/stopwatch
```

## Examples

### Basics

```go
// create a new stopwatch, the timer starts immediately.
s := stopwatch.Start()

// get elapsed duration at any time
duration := s.ElapsedTime()
// some work ... another elasped time
duration2 := s.ElapsedTime()

// create a new stopwatch, but do not start immediately
s := stopwatch.New()

// ... start it later
s.Start()
```

### Resume/Stop

```go
// reset the stopwatch
s.Reset()
// .. or stop the stopwatch
s.Stop()

// resume the timer after a reset/stop
s.Start()
```

### Lap

```go
// create a lap
lap1 := s.Lap()
lap2 := s.Lap()
lap3 := s.Lap()

// get a list of all lap durations
list := s.Laps()

// lap returns zero duration if the timer is stopped/resetted
s.Stop()
lap4 := s.Lap() // lap4 == time.Duration(0)
```

### Helpers
```go
// string representation of stopwatch
fmt.Printf("stopwatch: %s", s)
```

## Credits

 * [Fatih Arslan](https://github.com/fatih)

## License

The MIT License (MIT) - see LICENSE.md for more details
