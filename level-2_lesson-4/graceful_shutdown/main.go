package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("worker goes down")
				wg.Done()
				return
			default:
				fmt.Println("working...")
			}
		}
	}(ctx)

	select {
	case <-termChan:
		fmt.Println("graceful shutdown...")
		<-time.NewTicker(3 * time.Second).C
		cancelFunc()
	}

	wg.Wait()

}
