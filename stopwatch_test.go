package stopwatch

import (
	"strconv"
	"testing"
	"time"
)

func TestStopwatch(t *testing.T) {
	sw := New()

	if sw == nil {
		t.Error("New returns a nil struct")
	}

	if sw.start.IsZero() {
		t.Error("Start time returns a zero time.Time type")
	}
}

func TestStopwatch_ElapsedTime(t *testing.T) {
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

func TestStopwatch_Start(t *testing.T) {
	sw := New()
	time.Sleep(time.Millisecond * 300)
	sw.Stop()
	time.Sleep(time.Millisecond * 600) // should be not counted
	sw.Start()
	time.Sleep(time.Millisecond * 300)

	ms := int(RoundFloat(float64(sw.ElapsedTime()/time.Millisecond), 0))
	if ms != 600 {
		t.Errorf("Start: got: %d expected: %d\n", ms, 600)
	}
}

func TestStopwatch_Stop(t *testing.T) {
	sw := New()
	time.Sleep(time.Millisecond * 300)
	sw.Stop()
	time.Sleep(time.Millisecond * 200)

	ms := int(RoundFloat(float64(sw.ElapsedTime()/time.Millisecond), 0))
	if ms != 300 {
		t.Errorf("Stop: got: %d expected: %d\n", ms, 300)
	}
}

func TestStopwatch_Reset(t *testing.T) {
	sw := New()
	sw.Reset()

	if !sw.start.IsZero() {
		t.Error("Reset should reset the initial start timer")
	}

	sw.Start()
	time.Sleep(time.Millisecond * 400)
	ms := int(RoundFloat(float64(sw.ElapsedTime()/time.Millisecond), 0))
	if ms != 400 {
		t.Errorf("Reset: got: %d expected: %d\n", ms, 400)
	}
}

func TestStopwatch_Lap(t *testing.T) {}

// return rounded version of x with prec precision.
func RoundFloat(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}
