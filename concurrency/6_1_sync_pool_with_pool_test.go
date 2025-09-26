package concurrency

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

var bufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func workerWithPool(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	buf.WriteString(fmt.Sprintf("worker %d processed data", id))
	_ = buf.String()
	bufPool.Put(buf)
}

// Benchmark: with Pool
func BenchmarkWorkersWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		numWorkers := 100
		wg.Add(numWorkers)
		for j := 0; j < numWorkers; j++ {
			go workerWithPool(j, &wg)
		}
		wg.Wait()
	}
}

/*
Thread-safe mechanisms
→ Mỗi goroutine sẽ truy cập vào local pool riêng, dựa trên processor ID (P) hiện tại
Việc sử dụng sync.Pool trong trường hợp này đáng giá vì:
Tiết kiệm đáng kể bộ nhớ và giảm áp lực lên garbage collector
Thời gian chậm hơn chỉ là ~1.3μs (rất nhỏ)
Trong môi trường production với nhiều goroutines,
việc giảm memory allocation sẽ có tác động tích cực lớn hơn
*/
