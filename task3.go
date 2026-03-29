package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func startServer(ctx context.Context, name string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Duration(rand.Intn(500)) * time.Millisecond):
				out <- fmt.Sprintf("[%s] metric: %d", name, rand.Intn(100))
			}
		}
	}()
	return out
}

func FanIn(ctx context.Context, channels ...<-chan string) <-chan string {
	result := make(chan string)
	var wg sync.WaitGroup

	multiplex := func(c <-chan string) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-c:
				if !ok {
					return
				}
				result <- val
			}
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go multiplex(ch)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) [cite: 180]
	defer cancel()

	ch1 := startServer(ctx, "Alpha") [cite: 183]
	ch2 := startServer(ctx, "Beta") [cite: 185]
	ch3 := startServer(ctx, "Gamma") [cite: 187]

	ch4 := FanIn(ctx, ch1, ch2, ch3) [cite: 189]

	for val := range ch4 {
		fmt.Println(val) [cite: 192]
	}
}