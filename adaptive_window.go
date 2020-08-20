package adwin

type AdaptiveWindow interface {
	Add(float64)
	Detected() bool
	Size() int
	Sum() float64
	Conservative(bool)
}
