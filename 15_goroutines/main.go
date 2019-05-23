package main

import (
	"fmt"
	"sync"
)

// syncs multiple goroutines
var wg = sync.WaitGroup{}

func main() {
	var name = "Walter White"
	wg.Add(1) // tell waitgroup we want to sync this goroutine
	go func(n string) {
		fmt.Println(n)
		wg.Done() // this goroutine has done its execution
	}(name)
	name = "Willy Wonka"
	wg.Wait() // no sleep needed, just wait for the goroutines
	// we print Walter White, as we want
}
