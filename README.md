# Adwin and Adwin2

Adwin is a adaptive windowing algorithm. It is from `Learning from time-changing data with adaptive windowing, Bifet, Albert, and Ricard Gavalda; Proceedings of the 2007 SIAM international conference on data mining. Society for Industrial and Applied Mathematics, 2007`.

Adwin2 improves time and memory requirements by using Exponential Histograms.

See also:
- https://www.cs.upc.edu/~gavalda/DataStreamSeminar/files/Lecture6.pdf

## Usage

### Adwin

```go
	delta := 0.01
	adwin := NewAdwin(delta)
	adwin.Conservative(true) // if you wants

	// Add stream data
	adwin.Add(x)

	// Show status
	adwin.Size()
	adwin.Sum()
	adwin.Detected()
```

### Adwin2

```go
	delta := 0.01
	adwin2 := NewAdwin2(delta)
	adwin2.Conservative(true) // if you wants

	// Add stream data
	adwin2.Add(x)

	// Show status
	adwin2.Size()
	adwin2.Sum()
	adwin2.Detected()
```

## Installation

```sh
$ go get github.com/monochromegane/adwin
```

## License

[MIT](https://github.com/monochromegane/adwin/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)
