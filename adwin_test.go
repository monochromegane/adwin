package adwin

import (
	"testing"
)

func TestAdwinConservative(t *testing.T) {
	delta := 0.01
	adwin := NewAdwin(delta)
	adwin.Conservative(true)

	// Static
	for i := 0; i < 5; i++ {
		adwin.Add(10.0)
	}
	if size := adwin.Size(); size != 5 {
		t.Errorf("Size should be %d, but %d", 5, size)
	}
	if sum := adwin.Sum(); sum != 50.0 {
		t.Errorf("Size should be %f, but %f", 50.0, sum)
	}
	if detected := adwin.Detected(); detected {
		t.Errorf("Detected should be %t, but %t", false, detected)
	}

	// Non static
	x := 5.0
	adwin.Add(x)
	if size := adwin.Size(); size != 1 {
		t.Errorf("Size should be %d, but %d", 1, size)
	}
	if sum := adwin.Sum(); sum != 5.0 {
		t.Errorf("Size should be %f, but %f", 5.0, sum)
	}
	if detected := adwin.Detected(); !detected {
		t.Errorf("Detected should be %t, but %t", true, detected)
	}
}

func TestAdwin(t *testing.T) {
	delta := 0.01
	adwin := NewAdwin(delta)
	adwin.Conservative(false)

	// Static
	for i := 0; i < 50; i++ {
		adwin.Add(10.0)
	}
	if size := adwin.Size(); size != 50 {
		t.Errorf("Size should be %d, but %d", 50, size)
	}
	if sum := adwin.Sum(); sum != 500.0 {
		t.Errorf("Size should be %f, but %f", 500.0, sum)
	}
	if detected := adwin.Detected(); detected {
		t.Errorf("Detected should be %t, but %t", false, detected)
	}

	// Non static
	for i := 0; i < 10; i++ {
		adwin.Add(5.0)
	}
	if size := adwin.Size(); size != 18 {
		t.Errorf("Size should be %d, but %d", 18, size)
	}
	if sum := adwin.Sum(); sum != 130.0 {
		t.Errorf("Size should be %f, but %f", 130.0, sum)
	}
	if detected := adwin.Detected(); !detected {
		t.Errorf("Detected should be %t, but %t", true, detected)
	}
}
