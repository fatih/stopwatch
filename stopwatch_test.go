package stopwatch

import (
	"strconv"
	"testing"
	"time"
)

func TestStopWatch(t *testing.T) {
	sw := New()

	if sw == nil {
		t.Error("New returns a nil struct")
	}

	if sw.start.IsZero() {
		t.Error("Start time returns a zero time.Time type")
	}
}

func TestStopWatch_ElapsedTime(t *testing.T) {
	sw := New()

	elapsedDurations := make([]time.Duration, 0)

	for i := 1; i < 10; i++ {
		time.AfterFunc(time.Millisecond*100*time.Duration(i), func() {
			elapsedDurations = append(elapsedDurations, sw.ElapsedTime())
		})
	}

	time.Sleep(time.Second) // now collect all elapsed times

	// Better tests are welcome :)
	for i, elapsed := range elapsedDurations {
		ms := int(RoundFloat(float64(elapsed/time.Millisecond), 0))

		n := (i + 1) * 100
		if ms != n {
			t.Errorf("ElapsedTime: got: %d expected: %d\n", ms, n)

		}
	}

}

func TestStopWatch_Stop(t *testing.T)  {}
func TestStopWatch_Pause(t *testing.T) {}

func TestStopWatch_Reset(t *testing.T) {
	sw := New()
	sw.Reset()

	if !sw.start.IsZero() {
		t.Error("Reset should reset the initial start timer")
	}
}

func TestStopWatch_Lap(t *testing.T) {}

// return rounded version of x with prec precision.
func RoundFloat(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}
