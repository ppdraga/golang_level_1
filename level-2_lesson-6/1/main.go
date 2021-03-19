package main

// Написать программу, которая использует мьютекс для безопасного
// доступа к данным из нескольких потоков. Выполните трассировку программы

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int = 0
	fmt.Printf("counter = %v\n", counter)

	n := 5
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("counter = %v\n", counter)

}
