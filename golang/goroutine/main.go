package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()

			mu.Lock()
			n += i
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(n)
}
