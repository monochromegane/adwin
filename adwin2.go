package adwin

import "math"

type Adwin2 struct {
	delta  float64
	window *ExpHist
}

func NewAdwin2(delta float64) *Adwin2 {
	return &Adwin2{
		delta:  delta,
		window: NewExpHist(2),
	}
}

func (a *Adwin2) Add(x float64) {
	a.window.Add(x)

	detected := a.detectChanging()
	if detected {
		a.window.Drop()
	}
}

func (a *Adwin2) detectChanging() bool {
	n := a.window.Size()
	if n < 2 {
		return false
	}
	wSum := a.window.Sum()
	w1 := a.window.Tail()
	n1 := w1.size()
	w1Sum := w1.sum()
	n0 := n - n1
	w0Sum := wSum - w1Sum

	e := epsilonConservative(a.delta, n, n0, n1)
	uHatW0 := w0Sum / float64(n0)
	uHatW1 := w1Sum / float64(n1)
	return math.Abs(uHatW0-uHatW1) >= e
}
