package maps

import (
	"sync"
	"testing"
)

func BenchmarkThreadSafeMap(b *testing.B) {
	syncMap := NewThreadSafeMap()
	var wg sync.WaitGroup
	wg.Add(b.N * 2)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			configuration.Put("a", "b")
			wg.Done()
			configuration.Remove("a")
			wg.Done()
		}
	})

	wg.Wait()

}
