package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seed non deterministically?
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	y := rand.Intn(100)
	if x < y {
		fmt.Println("Hmm")
	} else {
		fmt.Println("mmm-hmmm")
	}
	fmt.Println(x, y)
	// now only assignment, previously,
	// we did initialize so we used :=
	x = rand.Intn(10)
	y = rand.Intn(10)

	if x < y {
		fmt.Println("wow")
	} else if x == y {
		fmt.Println("wowowow")
	} else {
		fmt.Println("what is this maaan")
	}
}
