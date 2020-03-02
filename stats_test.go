package adwin

import "testing"

func TestEpsilonConservative(t *testing.T) {
	delta := 0.5
	n := 8
	n0 := 2
	n1 := 6
	if result := epsilonConservative(delta, n, n0, n1); result > 1.1774 && 1.1775 < result {
		t.Errorf("EpsilonConservative should return about %f, but %f", 1.1774, result)
	}

}

func TestEpsilon(t *testing.T) {
	delta := 0.5
	v := 0.1
	n := 8
	n0 := 2
	n1 := 6
	if result := epsilon(v, delta, n, n0, n1); result > 1.4729 && 1.4730 < result {
		t.Errorf("Epsilon should return about %f, but %f", 1.4729, result)
	}

}

func TestVariance(t *testing.T) {
	w := []float64{19.0, 22.0, 22.0}
	if result := variance(w); result != 2.0 {
		t.Errorf("Variance should return %f, but %f", 2.0, result)
	}

}

func TestHarmonicMean(t *testing.T) {
	n0 := 2.0
	n1 := 6.0
	if result := harmonicMean(n0, n1); result != 1.5 {
		t.Errorf("HarmonicMean should return %f, but %f", 1.5, result)
	}
}

func TestSum(t *testing.T) {
	w := []float64{1.0, 2.0, 3.0}
	if result := sum(w); result != 6.0 {
		t.Errorf("Sum should return %f, but %f", 6.0, result)
	}
}

func TestAvarage(t *testing.T) {
	w := []float64{1.0, 2.0, 3.0}
	if result := average(w); result != 2.0 {
		t.Errorf("Average should return %f, but %f", 2.0, result)
	}
}
