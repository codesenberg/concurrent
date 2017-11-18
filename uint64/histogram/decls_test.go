package histogram

func testValues() []uint64 {
	return []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func testHistogram() *Histogram {
	return Default()
}

var benchValueSets = []struct {
	name string
	data []uint64
}{
	{"1-10", urange(1, 10)},
	{"1-100", urange(1, 100)},
	{"1-1000", urange(1, 1000)},
}
