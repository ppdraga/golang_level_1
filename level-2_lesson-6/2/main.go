package main

//Написать многопоточную программу, в которой будет использоваться явный вызов планировщика.
//Выполните трассировку программы

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("interrupt proc appears after 10! ")
	}()

	for i := 0; i < 20; i++ {
		fmt.Printf("%v, ", i)
		if i == 10 {
			runtime.Gosched()
		}
	}
	wg.Wait()

}
