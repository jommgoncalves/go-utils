package math

// Mean Struct to calculate the mean.
type Mean struct {
	Sum   int
	Count int
}

// NewMean Initializes struct.
func NewMean(sum, count int) *Mean {
	return &Mean{Sum: sum, Count: count}
}

// AddSum Increments sum.
func (m *Mean) AddSum(sum int) {
	m.Sum += sum
}

// AddCount Increments count.
func (m *Mean) AddCount(count int) {
	m.Count += count
}

// Calc Calculates the mean.
func (m *Mean) Calc() int {
	if m.Count == 0 {
		return 0
	}
	return Round(float64(m.Sum) / float64(m.Count))
}
