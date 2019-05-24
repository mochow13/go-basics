package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go goCount()
		go incCount()
	}
	wg.Wait()
}

/* counter will be printed in a very chaotic way and different
in each run. This is because the goroutines are racing against each
other to complete execution and there is no synchronization among them
*/

func goCount() {
	fmt.Println("Count is now ", counter)
	wg.Done()
}

func incCount() {
	counter++
	wg.Done()
}
