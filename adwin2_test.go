package adwin

import (
	"testing"
)

func TestAdwin2Conservative(t *testing.T) {
	delta := 0.01
	adwin := NewAdwin2(delta)
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
	for i := 0; i < 3; i++ {
		adwin.Add(5.0)
	}
	if size := adwin.Size(); size != 2 {
		t.Errorf("Size should be %d, but %d", 2, size)
	}
	if sum := adwin.Sum(); sum != 10.0 {
		t.Errorf("Size should be %f, but %f", 10.0, sum)
	}
	if detected := adwin.Detected(); !detected {
		t.Errorf("Detected should be %t, but %t", true, detected)
	}
}

func TestAdwin2(t *testing.T) {
	delta := 0.01
	adwin := NewAdwin2(delta)
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
	for i := 0; i < 15; i++ {
		adwin.Add(5.0)
	}
	if size := adwin.Size(); size != 17 {
		t.Errorf("Size should be %d, but %d", 17, size)
	}
	if sum := adwin.Sum(); sum != 95.0 {
		t.Errorf("Size should be %f, but %f", 95.0, sum)
	}
	if detected := adwin.Detected(); !detected {
		t.Errorf("Detected should be %t, but %t", true, detected)
	}
}
