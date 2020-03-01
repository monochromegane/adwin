package adwin

import (
	"math"
)

type Adwin struct {
	delta  float64
	window []float64
}

func NewAdwin(delta float64) *Adwin {
	return &Adwin{
		delta:  delta,
		window: []float64{},
	}
}

func (a *Adwin) Add(x float64) {
	a.window = append(a.window, x)
	for {
		detected := a.detectChanging()
		if detected {
			a.window = a.window[1:]
		} else {
			break
		}
	}
}

func (a *Adwin) detectChanging() bool {
	for i := 1; i < len(a.window); i++ {
		w0 := a.window[:i]
		w1 := a.window[i:]
		epsilon := epsilonConservative(a.delta, len(a.window), len(w0), len(w1))
		uHatW0 := average(w0)
		uHatW1 := average(w1)
		if math.Abs(uHatW0-uHatW1) >= epsilon {
			return true
		}
	}
	return false
}
