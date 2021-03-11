package main

import (
	"sync"
)

func main() {

	var mu sync.Mutex
	mu.Lock()
	unlockMutex(&mu)

}

func unlockMutex(m *sync.Mutex) {
	defer m.Unlock()
}
