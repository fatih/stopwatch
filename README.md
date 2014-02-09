# Stopwatch [![GoDoc](https://godoc.org/github.com/fatih/stopwatch?status.png)](http://godoc.org/github.com/fatih/stopwatch) [![Build Status](https://travis-ci.org/fatih/stopwatch.png)](https://travis-ci.org/fatih/stopwatch)

Stopwatch implements the common functions of a stop-watch.

For usage see examples below or click on the godoc badge.

## Install

```bash
go get github.com/fatih/stopwatch
```

## Example

```go
// create a set with zero items
sw := stopwatch.New()

// get elapsed duration at any time
duration := sw.ElapsedTime()

// reset the stopwatch
sw.Reset()

// string representation of stopwatch
fmt.Printf("stopwatch: %s", sw)

```

#### Basic Operations

## Credits

 * [Fatih Arslan](https://github.com/fatih)

## License

The MIT License (MIT) - see LICENSE.md for more details
