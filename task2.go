package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	// 1. Способ через sync/atomic
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println("Atomic Counter:", counter)

	// 2. Способ через sync.Mutex
	var countMutex int
	var mu sync.Mutex
	wg = sync.WaitGroup{} // сброс

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			countMutex++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Mutex Counter:", countMutex)
}