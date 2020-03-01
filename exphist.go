package adwin

type ExpHist struct {
	mergeSeize int
	buckets    buckets
}

func NewExpHist(m int) *ExpHist {
	return &ExpHist{
		mergeSeize: m + 1,
		buckets:    [][]bucket{[]bucket{}},
	}
}

func (e *ExpHist) Add(x float64) {
	bucket := bucket{
		content:  x,
		capacity: 1,
	}
	e.buckets[0] = append(e.buckets[0], bucket)

	e.merge()
}

func (e *ExpHist) Drop() {
	e.tail()
}

func (e *ExpHist) Size() int {
	return e.buckets.size()
}

func (e *ExpHist) Sum() float64 {
	return e.buckets.sum()
}

func (e *ExpHist) Tail() buckets {
	if len(e.buckets[len(e.buckets)-1]) == 1 {
		return [][]bucket{e.buckets[len(e.buckets)-1]}
	} else {
		return [][]bucket{[]bucket{e.buckets[len(e.buckets)-1][0]}}
	}
}

func (e *ExpHist) merge() {
	for i, _ := range e.buckets {
		if len(e.buckets[i]) < e.mergeSeize {
			continue
		}
		if i == len(e.buckets)-1 {
			e.buckets = append(e.buckets, []bucket{})
		}
		e.buckets[i+1] = append(e.buckets[i+1], bucket{
			content:  e.buckets[i][0].content + e.buckets[i][1].content,
			capacity: e.buckets[i][0].capacity + e.buckets[i][1].capacity,
		})
		e.buckets[i] = e.buckets[i][2:]
	}
}

func (e *ExpHist) tail() {
	if len(e.buckets[len(e.buckets)-1]) == 1 {
		e.buckets = e.buckets[0 : len(e.buckets)-1]
	} else {
		e.buckets[len(e.buckets)-1] = e.buckets[len(e.buckets)-1][1:]
	}
}

type buckets [][]bucket

func (b buckets) size() int {
	sum := 0
	for i, _ := range b {
		for j, _ := range b[i] {
			sum += b[i][j].capacity
		}
	}
	return sum
}

func (b buckets) sum() float64 {
	sum := 0.0
	for i, _ := range b {
		for j, _ := range b[i] {
			sum += b[i][j].content
		}
	}
	return sum
}

type bucket struct {
	content  float64
	capacity int
}
