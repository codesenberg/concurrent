package histogram

func hashString(s string) uint32 {
	// fnv32 hash
	const prime32 = 16777619
	hash := uint32(2166136261)
	for _, c := range []byte(s) {
		hash *= prime32
		hash ^= uint32(c)
	}
	return hash
}

// WithDefaultHash creates histogram with specified shardCount and
// reasonable sharding function.
func WithDefaultHash(shardsCount uint32) (*Histogram, error) {
	return New(shardsCount, hashString)
}

// Default creates histogram with reasonable defaults.
func Default() *Histogram {
	// We can safely ignore the error in this case
	h, _ := WithDefaultHash(32)
	return h
}
