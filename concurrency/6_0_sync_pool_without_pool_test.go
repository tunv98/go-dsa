package concurrency

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

func workerWithoutPool(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	buf := new(bytes.Buffer) // always create new buffer
	buf.WriteString(fmt.Sprintf("worker %d processed data", id))
	_ = buf.String()
}

// Benchmark: without Pool
func BenchmarkWorkersWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		numWorkers := 100
		wg.Add(numWorkers)
		for j := 0; j < numWorkers; j++ {
			go workerWithoutPool(j, &wg)
		}
		wg.Wait()
	}
}
