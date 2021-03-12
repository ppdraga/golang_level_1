package main

import (
	"fmt"
	"time"
)

func main() {

	var counter int = 0
	fmt.Printf("counter at start %v\n", counter)

	var workers = make(chan struct{}, 1)

	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}

		go func(job int) {
			defer func() {
				<-workers
			}()
			counter++
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Printf("counter at finish %v\n", counter)
}
