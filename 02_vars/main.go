package main

import "fmt"

// we can use const and var to declare global variables
// we can declare variables in a shorter form but only inside a function
const life = "hashtag"
const long int64 = 20398470213

var hashtag = "life"
var just int8 = 12

func declareSomeShortFormedVariables() {
	aname := "what"
	fmt.Println(aname)
	wow := 2.1441139 // so by default float 64
	fmt.Printf("%T\n", wow)
	// if you don't use a variable, it's an ERROR!
	// Lesson: write only that matters
	name, email := "mottakin", "mochow@gmail.com"
	// can do this in one line
	fmt.Println(name, email)
}

func main() {
	// can create variables in few ways
	// using var
	var name = "Comrades!"     // can avoid 'string' here
	var age = 1234533378933310 // can avoid writing int if it is of type int
	fmt.Println(name, age+7)
	fmt.Printf("%T\n", age)

	// size of `int` varies with machine
	// on my machine, its 64 bit

	var bro int32 = 12415               // but cannot avoid int32, it has to be explicit
	var coolbro = false                 // you guessed it right
	fmt.Printf("%T %T\n", bro, coolbro) // type of big, prints int32

	const saymyname = "Walter White" // constant
	declareSomeShortFormedVariables()
}
