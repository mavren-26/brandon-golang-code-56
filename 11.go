package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range tasks {
		result := n * n
		fmt.Printf("worker %d processed %d → %d\n", id, n, result)
		time.Sleep(200 * time.Millisecond) // simulate work
	}
}

func main() {
	const workerCount = 3
	tasks := make(chan int, 10)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Send tasks
	for i := 1; i <= 10; i++ {
		tasks <- i
	}
	close(tasks)

	// Wait for workers
	wg.Wait()
}
