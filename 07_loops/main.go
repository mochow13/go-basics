package main

import "fmt"

// only loop of Go is for loop
func main() {
	// fizzbuzz challenge! yayy!
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d FizzBuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d Buzz\n", i)
		}
	}

	// we can use it like while loops
	i := 1
	for i <= 10 {
		fmt.Println("Prity")
		i++
	}
}
