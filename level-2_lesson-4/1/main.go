package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// number of threads
	n := 100

	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(id int) {
			fmt.Printf("Start goroutine %v\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Stop goroutine %v\n", id)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
