package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Idea: A goroutine is a lightweight thread. If main returns, all goroutines end immediately.

func sequential() {
	hi("Tu")
}

func basicGoroutineWithSleep() {
	fmt.Println("[0] Start")
	go func() {
		hi("[0] Hello World")
	}()
	time.Sleep(10 * time.Millisecond) // wait for goroutine run
	fmt.Println("[0] Done")
}

/*
Add: increase the value of the counter by the specified amount
Done: decrease the value of the counter by one
Wait blocks until the value of the counter is zero
*/
func basicGoroutineWithWaitGroup() {
	fmt.Println("[1] Start")
	var wg sync.WaitGroup // create WaitGroup
	wg.Add(1)             // tell wg have 1 goroutine to wait for
	go func() {
		defer wg.Done() // mark done when goroutine finishes
		fmt.Println("[1] Hello there!")
	}()
	wg.Wait() // wait until counter is back to 0
	fmt.Println("[1] Done")
}

/*
Key points:
you can move wg.Add() out of the for loop, you need to be careful
-> the number you pass to Add must exactly match the total number of goroutines
Risk: if you miscount, your program can:
+ Hang forever (if you Add too much - Wait never returns)
+ Panic (if you `Done` more times than you `Add`) -> fatal error: all goroutines are asleep - deadlock!
+ Miscount (if you `Done` fewer times than you `Add`) -> no wait for last goroutine
*/
func multipleGoroutineWithWaitGroup() {
	fmt.Println("[2] Start")
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		//wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("[2] Hello %d\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("[2] Done")
}

func Test_concurrency(t *testing.T) {
	//sequential()
	//basicGoroutineWithSleep()
	//basicGoroutineWithWaitGroup()
	multipleGoroutineWithWaitGroup()
}
