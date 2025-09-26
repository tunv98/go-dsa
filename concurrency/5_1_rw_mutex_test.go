package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
sync.RWMutex
+ Có 2 loại khóa: đọc (RLock) và ghi (Lock).
+ RLock: cho phép nhiều goroutine đọc đồng thời, nhưng không cho phép ghi.
+ Lock: chỉ cho phép một goroutine ghi, và sẽ chặn tất cả các goroutine khác (bao gồm cả đọc).
Cho phép tối ưu hiệu năng khi đọc nhiều, ghi ít.
*/

type SafeMap struct {
	sm map[string]string
	mu sync.RWMutex
}

func newSafeMap() *SafeMap {
	return &SafeMap{
		sm: make(map[string]string),
	}
}

func (m *SafeMap) Get(key string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, existed := m.sm[key]
	time.Sleep(500 * time.Millisecond)
	return value, existed
}

func (m *SafeMap) Set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.sm[key] = value
	time.Sleep(1 * time.Second) // write sẽ block read
}

func Test_SafeMap(t *testing.T) {
	sm := newSafeMap()
	// Writer goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			sm.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("val%d", i)) // block read
			fmt.Println("Wrote key", i)
		}
	}()

	// Reader goroutines
	var wg sync.WaitGroup
	for r := 1; r <= 3; r++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i <= 10; i++ {
				if v, ok := sm.Get("key1"); ok { // có thể đọc đồng thời
					fmt.Printf("[Reader %d] Read: key1 = %s\n", id, v)
				} else {
					fmt.Printf("[Reader %d] key1 not found\n", id)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}(r)
	}
	wg.Wait()
}
