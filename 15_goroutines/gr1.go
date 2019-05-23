package main

import (
	"fmt"
	"time"
)

func main() {
	go sayMyName() // spinning up a green thread
	// without this, main would execute immediately, won't let sayMyName() to execute
	time.Sleep(100 * time.Millisecond)
}

func sayMyName() {
	fmt.Println("Walter White")
}
