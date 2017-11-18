package histogram

func urange(from, to uint64) []uint64 {
	if to < from {
		panic("to must be less than from")
	}
	result := make([]uint64, to-from+1)
	for i := from; i <= to; i++ {
		result[i-from] = i
	}
	return result
}
