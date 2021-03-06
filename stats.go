package adwin

import "math"

func epsilonConservative(delta float64, n, n0, n1 int) float64 {
	m := harmonicMean(float64(n0), float64(n1))
	deltaDash := delta / float64(n)
	return math.Sqrt(1.0 / (2.0 * m) * math.Log(4.0/deltaDash))
}

func epsilon(v, delta float64, n, n0, n1 int) float64 {
	m := harmonicMean(float64(n0), float64(n1))
	deltaDash := delta / math.Log(float64(n))
	l := math.Log(2.0 / deltaDash)
	return math.Sqrt((2.0/m)*v*l) + (2.0/(3.0*m))*l
}

func variance(w []float64) float64 {
	sum := 0.0
	sum2 := 0.0
	for i, _ := range w {
		sum += w[i]
		sum2 += w[i] * w[i]
	}
	n := float64(len(w))
	avg := sum / n
	return sum2/n - avg*avg
}

func harmonicMean(n0, n1 float64) float64 {
	return 1.0 / ((1.0 / n0) + (1.0 / n1))
}

func sum(w []float64) float64 {
	sum := 0.0
	for _, v := range w {
		sum += v
	}
	return sum
}

func average(w []float64) float64 {
	return sum(w) / float64(len(w))
}
