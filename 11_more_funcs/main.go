package main

import "fmt"

// takes a function as an argument
// the argument function takes two inputs and returns an integer
func operate(xor func(x, y int) int) {
	fmt.Println("xor val:", xor(15, 23))
}

// returns a function with arguments x,y int and return value int
func getFunc() func(int, int) int {
	// we are returning `ret`
	ret := func(x, y int) int {
		return x ^ y
	}
	return ret
}

func grow() func(string) string {
	var species string
	species = ""
	f := func(y string) string {
		species = species + y
		return species
	}
	return f
}

// Go has anonymous functions and closures
func main() {
	// assign functions to variables
	fun := func(x int) {
		fmt.Println("Wow this is assignable!")
		fmt.Println("with an argument", x)
	}
	// invoke it
	fun(12)
	// function invocation while declaring
	func(x int) {
		fmt.Println("This is being invoked while begin declared!")
		fmt.Println("with value", x)
	}(42)

	// pass functions to higher order ones
	lambda := func(x, y int) int { // lambda is a variable
		return x ^ y
	}
	operate(lambda)

	// returning function
	getop := getFunc()
	fmt.Println("getop:", getop(2, 12))

	// closures
	a := grow()
	b := grow()

	fmt.Println(a("I"))
	fmt.Println(b("Cornerstone"))
	fmt.Println(a(" won't ever be your")) // outputs "I won't ever be your"
}
