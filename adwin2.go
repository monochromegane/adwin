package adwin

import (
	"math"

	exphist "github.com/monochromegane/exponential-histograms"
)

type Adwin2 struct {
	delta   float64
	window  *exphist.ExpHistRealNumber
	window2 *exphist.ExpHistRealNumber

	conservative bool
	detected     bool
}

func NewAdwin2(delta float64) AdaptiveWindow {
	return &Adwin2{
		delta:   delta,
		window:  exphist.NewForRealNumber(2),
		window2: exphist.NewForRealNumber(2),
	}
}

func (a *Adwin2) Conservative(t bool) {
	a.conservative = t
}

func (a *Adwin2) Add(x float64) {
	a.window.Add(x)
	if !a.conservative {
		a.window2.Add(x * x)
	}
	a.detected = false

	detected := a.detectChanging()
	if detected {
		a.window.Drop()
		if !a.conservative {
			a.window2.Drop()
		}
		a.detected = true
	}
}

func (a *Adwin2) Size() int {
	return a.window.Size()
}

func (a *Adwin2) Sum() float64 {
	return a.window.Sum()
}

func (a *Adwin2) Detected() bool {
	return a.detected
}

func (a *Adwin2) detectChanging() bool {
	n := a.window.Size()
	if n < 2 {
		return false
	}
	wSum := a.window.Sum()
	w1 := a.window.Tail()
	n1 := w1.Size()
	w1Sum := w1.Sum()
	n0 := n - n1
	w0Sum := wSum - w1Sum

	e := 0.0
	if a.conservative {
		e = epsilonConservative(a.delta, n, n0, n1)
	} else {
		v := a.window2.Sum()/float64(a.window2.Size()) - ((wSum / float64(n)) * (wSum / float64(n)))
		e = epsilon(v, a.delta, n, n0, n1)
	}
	uHatW0 := w0Sum / float64(n0)
	uHatW1 := w1Sum / float64(n1)
	return math.Abs(uHatW0-uHatW1) >= e
}
