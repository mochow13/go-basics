package main

import (
	"fmt"
	"time"
)

func main() {
	var name = "Walter White"
	/* anonymous function, invoking immediately
	the function won't print Walter White because before
	it gets a chance, the name is reassigned, so it prints
	Willy Wonka */
	go func() {
		fmt.Println(name)
	}()
	name = "Willy Wonka"
	/* sleep is still necessary because main executes before even
	the anonymous function has a scope to print the name */
	time.Sleep(100 * time.Millisecond)
}
