package concurrency

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

/*
sync.Mutex
Lock: chỉ 1 goroutine được vào critical section.
Unlock: giải phóng để goroutine khác vào.
Đọc và ghi đều phải dùng Lock/Unlock.
Không phân biệt giữa đọc và ghi — mọi truy cập đều bị chặn nếu có goroutine đang giữ khóa.
*/
type Counter struct {
	value int
	mu    sync.Mutex
}

func newCounter(value int) *Counter {
	return &Counter{
		value: value,
		mu:    sync.Mutex{},
	}
}
func (c *Counter) Inc(val int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += val
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	v := c.value
	return v
}

/*
wg: biết cần chờ 2 goroutine.

start: là barrier channel(kênh trống, ko buffer) để ép 2 goroutine bắt đầu gần như 1 lúc
-> tăng khả năng xảy ra tranh chấp lock

close(start): đóng channel start làm cho tất cả goroutine đang chờ <-start được giải phóng ngay lập tức
-> ko gửi gì cả -> tiếp tục chạy hàm sau

done := make(chan struct{})
go func() { wg.Wait(); close(done) }()
select {
case <-done:

	// ok, cả hai đã xong

case <-time.After(2 * time.Second):

	    t.Fatal("timeout waiting for goroutines to finish")
	}

-> wg.Wait() không hỗ trợ select/timeout trực tiếp,
nên ta bọc nó trong một goroutine và phát tín hiệu qua done.
*/
func Test_counter(t *testing.T) {
	// want to test concurrency here
	counter := newCounter(1)
	const goroutines = 2
	var wg sync.WaitGroup
	wg.Add(goroutines)

	start := make(chan struct{}) //barrier
	go func() {
		defer wg.Done()
		<-start // block -> đợi close
		counter.Inc(3)
	}()
	go func() {
		defer wg.Done()
		<-start // block -> đợi close
		counter.Inc(2)
	}()
	close(start)

	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
		fmt.Println("done")
	case <-time.After(2 * time.Second):
		t.Fatal("timeout waiting for goroutines to finish")
	}

	got := counter.Value()
	expected := 1 + 3 + 5
	assert.NotEqualf(t, got, expected, fmt.Sprintf("expected %d, but receive %d", expected, got))
}

func Test_counter_manyGoroutine(t *testing.T) {
	counter := newCounter(0)

	const goroutines = 1000
	var wg sync.WaitGroup
	wg.Add(goroutines)
	start := make(chan struct{})
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			<-start
			counter.Inc(1)
		}()
	}
	close(start)
	wg.Wait()
	if got := counter.Value(); got != goroutines {
		t.Fatalf("final value = %d, want %d", got, goroutines)
	}
}
