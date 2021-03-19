package main

import (
	"fmt"
	"sync"
)

//Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”

func main() {

	var wg sync.WaitGroup
	var counter int = 1

	n := 5
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			counter++
			fmt.Println(counter)
		}()
	}
	wg.Wait()

}

/*
go run -race main.go

2
3
4
==================
WARNING: DATA RACE
Read at 0x00c0000180d0 by goroutine 8:
main.main.func1()
/home/yaroslav/go/src/github.com/ppdraga/golang_level_1/level-2_lesson-6/3/main.go:21 +0x8f

Previous write at 0x00c0000180d0 by goroutine 7:
main.main.func1()
/home/yaroslav/go/src/github.com/ppdraga/golang_level_1/level-2_lesson-6/3/main.go:21 +0xa8

Goroutine 8 (running) created at:
main.main()
/home/yaroslav/go/src/github.com/ppdraga/golang_level_1/level-2_lesson-6/3/main.go:19 +0xe9

Goroutine 7 (finished) created at:
main.main()
/home/yaroslav/go/src/github.com/ppdraga/golang_level_1/level-2_lesson-6/3/main.go:19 +0xe9
==================
6
5
Found 1 data race(s)
exit status 66
*/
