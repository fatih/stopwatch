# Stopwatch [![GoDoc](https://godoc.org/github.com/fatih/stopwatch?status.svg)](http://godoc.org/github.com/fatih/stopwatch) [![Build Status](https://travis-ci.org/fatih/stopwatch.svg)](https://travis-ci.org/fatih/stopwatch)

Stopwatch implements a simple stopwatch functionality. Features :

* Two possible ways to create a Stopwatch. Initialized or uninitialized.
* Start/Stop at any time or Reset.
* Take an individual Lap time
* Stores the list of each Lap
* Satisfies JSON Marshaler/Unmarshaler interface
* Handy methods like Print()/Log() to log a function execution time with one step.

Feel free to fork and send a pull request for any
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
s := stopwatch.Start(0)

// get elapsed duration at any time
duration := s.ElapsedTime()
// some work ... another elapsed time
duration2 := s.ElapsedTime()

// create a new stopwatch, but do not start immediately
s := stopwatch.New()
// ... start it later
s.Start()

// start a stopwatch after a certain time
s := stopwatch.Start(2 * time.Second)
d1 := s.ElapsedTime() // d1 is zero here
// after two seconds it works
d2 := s.ElapsedTime()
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

// lap returns zero duration if the timer is stopped/reseted
s.Stop()
lap4 := s.Lap() // lap4 == time.Duration(0)
```

### Helpers
```go
// String representation of stopwatch
fmt.Printf("stopwatch: %s", s)

// find out how long a function lasts
// outputs when the function returns:  myFunction - elapsed: 2.000629842s
defer Start(0).Print("myfunction")

// Marshal to a JSON object.
type API struct {
    Name      string     `json:"name"`
    Stopwatch *Stopwatch `json:"elapsed"`
}

a := API{
    Name:      "Example API Call",
    Stopwatch: Start(0),
}

// do some work ...
time.Sleep(time.Millisecond * 20)

b, err := json.Marshal(a)
if err != nil {
    t.Errorf("error: %s\n", err)
}

// output: {"name":"Example API Call","elapsed":"21.351657ms"}
fmt.Println(string(b))

// Unmarshal from a JSON object.
v := new(API)
err = json.Unmarshal(b, v)
if err != nil {
    t.Errorf("error: %s\n", err)
}

// Get back our elapsed time
duration := v.Stopwatch.ElapsedTime()
```

## Credits

 * [Fatih Arslan](https://github.com/fatih)

## License

The MIT License (MIT) - see LICENSE.md for more details
