package histogram

func testValues() []float64 {
	return []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func testHistogram() *Histogram {
	return Default()
}

const distributionSize = 10000

var benchValueSets = []struct {
	name string
	data []float64
}{
	{"uniform", uniform(0, 10000, distributionSize)},
	{"normal", normal(5000, 500, distributionSize)},
}
