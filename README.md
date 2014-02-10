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
sw := stopwatch.New()

// get elapsed duration at any time
duration := sw.ElapsedTime()

// some work ...
duration2 := sw.ElapsedTime()
```

### Resume/Stop

```go
// reset the stopwatch
sw.Reset()
// .. or stop the stopwatch
sw.Stop()

// resume the timer after a reset/stop
sw.Start()
```

### Lap

```go
// create a lap
lap1 := sw.Lap()
lap2 := sw.Lap()
lap3 := sw.Lap()

// get a list of all lap durations
list := sw.Laps()

// lap returns zero duration if the timer is stopped/resetted
sw.Stop()

lap4 := sw.Lap() // lap4 == time.Duration(0)
```

### Helpers
```go
// string representation of stopwatch
fmt.Printf("stopwatch: %s", sw)
```

## Credits

 * [Fatih Arslan](https://github.com/fatih)

## License

The MIT License (MIT) - see LICENSE.md for more details
