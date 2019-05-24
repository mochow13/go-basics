package main

import (
	"fmt"
	"sync"
)

// concurrency and parallelism is destroyed in this example
// it is like a sequential execution, but we show this as an example of mutex
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		// Lock on goroutine, not inside the goCount(), but outside
		// doing a read lock
		m.RLock()
		go goCount()
		m.Lock()
		go incCount()
	}
	wg.Wait()
}

func goCount() {
	fmt.Println("Count is now ", counter)
	m.RUnlock() // unlock when we are done with operation
	wg.Done()
}

func incCount() {
	counter++
	m.Unlock()
	wg.Done()
}
