package adwin

import (
	"math"
)

type Adwin struct {
	delta  float64
	window []float64

	conservative bool
	detected     bool
}

func NewAdwin(delta float64) *Adwin {
	return &Adwin{
		delta:  delta,
		window: []float64{},
	}
}

func (a *Adwin) Conservative(t bool) {
	a.conservative = t
}

func (a *Adwin) Add(x float64) {
	a.window = append(a.window, x)
	a.detected = false
	for {
		detected := a.detectChanging()
		if detected {
			a.detected = true
			a.window = a.window[1:]
		} else {
			break
		}
	}
}

func (a *Adwin) Size() int {
	return len(a.window)
}

func (a *Adwin) Sum() float64 {
	return sum(a.window)
}

func (a *Adwin) Detected() bool {
	return a.detected
}

func (a *Adwin) detectChanging() bool {
	v := 0.0
	if !a.conservative {
		v = variance(a.window)
	}

	for i := 1; i < len(a.window); i++ {
		w0 := a.window[:i]
		w1 := a.window[i:]
		e := 0.0
		if a.conservative {
			e = epsilonConservative(a.delta, len(a.window), len(w0), len(w1))
		} else {
			e = epsilon(v, a.delta, len(a.window), len(w0), len(w1))
		}
		uHatW0 := average(w0)
		uHatW1 := average(w1)
		if math.Abs(uHatW0-uHatW1) >= e {
			return true
		}
	}
	return false
}
