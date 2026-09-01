package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mu sync.Mutex

func workerSlow(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker %d start\n", id)

	mu.Lock()
	time.Sleep(10 * time.Millisecond)
	count++
	mu.Unlock()
	fmt.Printf("Worker A-%d done, count=%d\n", id, count)

}

func workerFast(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker %d start\n", id)

	time.Sleep(10 * time.Millisecond)
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Printf("Worker A-%d done, count  =%d\n", id, count)
}
