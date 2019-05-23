package main

import (
	"fmt"
	"time"
)

func main() {
	var name = "Walter White"
	// pass name before it gets changed, a way to solve the problem
	go func(n string) {
		fmt.Println(n)
	}(name)
	name = "Willy Wonka"
	/* sleep is still necessary because main executes before even
	the anonymous function has a scope to print the name */
	time.Sleep(100 * time.Millisecond)
}
