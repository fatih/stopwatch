package stopwatch

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"
)

func TestStopwatchNew(t *testing.T) {
	sw := New()

	if sw == nil {
		t.Error("New returns a nil struct")
	}

	if !sw.start.IsZero() {
		t.Error("New time returns a non zero time.Time type")
	}

	if sw.laps == nil {
		t.Error("New should initialize the laps array")
	}
}

func TestStopwatchStart(t *testing.T) {
	sw := Start(0)

	if sw == nil {
		t.Error("Start returns a nil struct")
	}

	if sw.start.IsZero() {
		t.Error("Start time returns a zero time.Time type")
	}

	if sw.laps == nil {
		t.Error("Start should initialize the laps array")
	}
}

func TestStopwatchAfter(t *testing.T) {
	a := Start(time.Millisecond * 30)

	if a == nil {
		t.Error("After returns a nil struct")
	}

	if !a.start.After(time.Now()) {
		t.Error("After should create a future based start time.")
	}

	time.Sleep(time.Millisecond * 10)
	ms1 := int(RoundFloat(float64(a.ElapsedTime()/time.Millisecond), 0))
	if ms1 != -20 {
		t.Errorf("After: got: %d expected: 0\n", ms1)
	}

	time.Sleep(time.Millisecond * 30)
	ms2 := int(RoundFloat(float64(a.ElapsedTime()/time.Millisecond), 0))
	if ms2 != 10 {
		t.Errorf("After: got: %d expected: %d\n", ms2, 20)
	}

}

func TestStopwatch_ElapsedTime(t *testing.T) {
	sw := Start(0)

	elapsedDurations := make([]time.Duration, 0)

	for i := 1; i < 10; i++ {
		time.AfterFunc(time.Millisecond*10*time.Duration(i), func() {
			elapsedDurations = append(elapsedDurations, sw.ElapsedTime())
		})
	}

	time.Sleep(time.Millisecond * 100) // now collect all elapsed times

	// Better tests are welcome :)
	for i, elapsed := range elapsedDurations {
		ms := int(RoundFloat(float64(elapsed/time.Millisecond), 0))

		n := (i + 1) * 10
		if ms != n {
			t.Errorf("ElapsedTime: got: %d expected: %d\n", ms, n)
		}
	}

	n := New()
	nl := n.ElapsedTime()
	if nl != time.Duration(0) {
		t.Errorf("ElapsedTime: got: %d, expected: 0\n", nl)
	}

}

func TestStopwatch_Start(t *testing.T) {
	sw := Start(0)
	time.Sleep(time.Millisecond * 30)
	sw.Stop()
	time.Sleep(time.Millisecond * 60) // should be not counted
	sw.Start(0)
	time.Sleep(time.Millisecond * 30)

	ms := int(RoundFloat(float64(sw.ElapsedTime()/time.Millisecond), 0))
	if ms != 60 {
		t.Errorf("Start: got: %d expected: %d\n", ms, 60)
	}
}

func TestStopwatch_Stop(t *testing.T) {
	sw := Start(0)
	time.Sleep(time.Millisecond * 30)
	sw.Stop()
	time.Sleep(time.Millisecond * 20)

	ms := int(RoundFloat(float64(sw.ElapsedTime()/time.Millisecond), 0))
	if ms != 30 {
		t.Errorf("Stop: got: %d expected: %d\n", ms, 30)
	}
}

func TestStopwatch_Reset(t *testing.T) {
	sw := Start(0)
	sw.Reset()

	if !sw.start.IsZero() {
		t.Error("Reset should reset the initial start timer")
	}

	sw.Start(0)
	time.Sleep(time.Millisecond * 40)
	ms := int(RoundFloat(float64(sw.ElapsedTime()/time.Millisecond), 0))
	if ms != 40 {
		t.Errorf("Reset: got: %d expected: %d\n", ms, 40)
	}
}

func TestStopwatch_Lap(t *testing.T) {
	sw := Start(0)

	time.Sleep(time.Millisecond * 10)
	lap1 := sw.Lap()

	time.Sleep(time.Millisecond * 20)
	lap2 := sw.Lap()

	time.Sleep(time.Millisecond * 30)
	lap3 := sw.Lap()

	ms1 := int(RoundFloat(float64(lap1/time.Millisecond), 0))
	ms2 := int(RoundFloat(float64(lap2/time.Millisecond), 0))
	ms3 := int(RoundFloat(float64(lap3/time.Millisecond), 0))

	if ms1 != 10 && ms2 != 20 && ms3 != 30 {
		t.Errorf("Lap: got: %d %d %d, expecting: %d %d %d\n",
			ms1, ms2, ms3, 10, 20, 30)
	}

	if len(sw.laps) != 3 {
		t.Error("Lap: number of laps should be 3")
	}

	sw.Stop()
	l := sw.Lap()
	if l != time.Duration(0) {
		t.Errorf("Lap: stopwatch is stopped but lap returns %d\n", l)
	}

	n := Start(0)
	n.Reset()

	u := n.Lap()
	if u != time.Duration(0) {
		t.Errorf("Lap: stopwatch is resetted but lap returns %d\n", u)
	}

}

func TestStopwatch_JSON(t *testing.T) {
	type API struct {
		Name      string     `json:"name"`
		Stopwatch *Stopwatch `json:"elapsed"`
	}

	a := API{
		Name:      "testData",
		Stopwatch: Start(0),
	}

	time.Sleep(time.Millisecond * 20)

	b, err := json.Marshal(a)
	if err != nil {
		t.Errorf("error: %s\n", err)
	}

	v := new(API)
	err = json.Unmarshal(b, v)
	if err != nil {
		t.Errorf("error: %s\n", err)
	}

	ms := int(RoundFloat(float64(v.Stopwatch.ElapsedTime()/time.Millisecond), 0))
	if ms != 20 {
		t.Errorf("json: got: %d, expected: %d", ms, 20)
	}

}

// return rounded version of x with prec precision.
func RoundFloat(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}
