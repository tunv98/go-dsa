package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Why worker pool?
when you have lots of tasks but want to limit how many run at the same time (to control CPU, memory, API rate, etc) -> a worker pool is perfect
Fan-out: send tasks to multiple workers (goroutines)
Fan-in: collect all results back into a single place
*/

/*
Close(channel) -> tell the receivers: "no more values will ever be sent to this channel"
*/

/*
read-only channel: `jobs <-chan int`
write-only channel: `results chan <- int`
bidirectional channel: `job chan int`
*/

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker %d started job %d \n", id, job)
		time.Sleep(500 * time.Millisecond) // process work
		fmt.Printf("worker %d finished job %d \n", id, job)
		results <- job * 2
	}
}

func fanInOut() {
	numJobs := 5
	numsWorkers := 2

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	//start workers(fan‑out)
	for w := 1; w <= numsWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	//enqueue jobs
	for j := 0; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) //tells workers “no more tasks”.

	go func() {
		wg.Wait()
		close(results) // after all workers finish sending
	}()
	for r := range results {
		fmt.Println("result:", r)
	}
}

func Test_fanInOut(t *testing.T) {
	fanInOut()
}
