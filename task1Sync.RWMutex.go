package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.RWMutex 
	safeMap := make(map[string]int)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			mu.Lock()
			safeMap["key"] = key
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	
	mu.RLock()
	fmt.Printf("Value (Mutex): %d\n", safeMap["key"]) 
	mu.RUnlock() 
}