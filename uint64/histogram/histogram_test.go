package histogram

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestHistogramIncrement(t *testing.T) {
	h := testHistogram()
	key := testValues()[0]
	h.Increment(key)
	if c := h.Get(key); c != 1 {
		t.Errorf("expected to get 1, but got %v", c)
	}
}

func TestHistogramCount(t *testing.T) {
	h := testHistogram()
	keys := testValues()
	for _, k := range keys {
		h.Increment(k)
	}
	if a, e := h.Count(), uint64(len(keys)); e != a {
		t.Errorf("expected %v entries, but found %v", e, a)
	}
}

func TestHistogramVisitAll(t *testing.T) {
	h := testHistogram()
	keys := testValues()
	keysMap := make(map[uint64]struct{})
	for _, k := range keys {
		h.Increment(k)
		keysMap[k] = struct{}{}
	}
	foundKeys := make(map[uint64]struct{})
	h.VisitAll(func(key uint64, c uint64) bool {
		if c != 1 {
			t.Errorf("expected %v for key %v, but got %v", 1, key, c)
			return false
		}
		foundKeys[key] = struct{}{}
		return true
	})
	for k := range foundKeys {
		delete(keysMap, k)
	}
	if len(keysMap) != 0 {
		t.Log("following keys where not found in histogram")
		for k := range keysMap {
			t.Log(k)
		}
		t.Fail()
	}
}

func TestHistogramConcurrentIncrement(t *testing.T) {
	keys := testValues()
	counters := make(map[uint64]*uint64)
	for _, k := range keys {
		counters[k] = new(uint64)
	}
	goroutinesCount := 64
	doneCh := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(goroutinesCount)
	h := testHistogram()
	for i := 0; i < goroutinesCount; i++ {
		go func() {
			defer wg.Done()
			idx := 0
			for {
				select {
				case <-doneCh:
					return
				default:
					key := keys[idx]
					h.Increment(key)
					atomic.AddUint64(counters[key], 1)
					idx = (idx + 1) % len(keys)
				}
			}
		}()
	}
	time.Sleep(100 * time.Millisecond)
	close(doneCh)
	wg.Wait()
	for k, exp := range counters {
		act := h.Get(k)
		if *exp != act {
			t.Errorf("for %v - got %v, want %v", k, act, exp)
			break
		}
	}
}

func BenchmarkHistogramPerformance(b *testing.B) {
	parallelismLevels := []int{1, 2, 4, 8, 16, 32}
	for _, p := range parallelismLevels {
		for _, set := range benchValueSets {
			b.Run(
				fmt.Sprintf(
					"goroutines=GOMAXPROCS*%v,set=%v", p, set.name,
				),
				benchmarkHistogramPerformance(p, set.data),
			)
		}
	}
}

func benchmarkHistogramPerformance(p int, data []uint64) func(*testing.B) {
	return func(b *testing.B) {
		h := testHistogram()
		b.SetParallelism(p)
		dlen := uint64(len(data))
		b.RunParallel(func(pb *testing.PB) {
			idx := uint64(0)
			key := data[idx]
			for pb.Next() {
				h.Increment(key)
				idx = atomic.AddUint64(&idx, 1) % dlen
				key = data[idx]
			}
		})
	}
}
