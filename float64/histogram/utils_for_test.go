package histogram

import "math/rand"

func uniform(min, max, size int) []float64 {
	result := make([]float64, size)
	diff := float64(max - min)
	for i := range result {
		result[i] = float64(min) + rand.Float64()*diff
	}
	return result
}

func normal(mean, stdev, size int) []float64 {
	result := make([]float64, size)
	for i := range result {
		result[i] = rand.NormFloat64()*float64(stdev) + float64(mean)
	}
	return result
}
