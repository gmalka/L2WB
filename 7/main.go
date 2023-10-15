package main

import (
	"fmt"
	"time"
)

func main() {
	// Пример использования функции:
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(500*time.Millisecond),
		sig(500*time.Millisecond),
		sig(500*time.Millisecond),
		sig(500*time.Millisecond),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
}

func or(channels ...<-chan any) <-chan any {
	end := make(chan struct{})
	done := make(chan any)

	for _, v := range channels {
		go func(v <-chan any) {
			select {
			case <-v:
				close(end)
				close(done)
				return
			case <-end:
				return
			}
		}(v)
	}

	return done
}