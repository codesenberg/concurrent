package histogram

func testValues() []string {
	return []string{"a", "b", "c", "aa", "ab", "ac",
		"ba", "bb", "bc", "ca", "cb", "cc"}
}

func testHistogram() *Histogram {
	return Default()
}

var benchValueSets = []struct {
	name string
	data []string
}{
	{"small-set-of-strings", testValues()},
	{"few-of-small-strings", generateStrings(172831233, 100, 10)},
	{"few-of-medium-strings", generateStrings(172831233, 100, 100)},
	{"few-of-big-strings", generateStrings(172831233, 100, 1000)},
	{"lots-of-small-strings", generateStrings(172831233, 10000, 10)},
	{"lots-of-medium-strings", generateStrings(172831233, 10000, 100)},
	{"lots-of-big-strings", generateStrings(172831233, 10000, 1000)},
}
