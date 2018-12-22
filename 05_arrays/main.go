package main

import "fmt"

func main() {
	// declare then assign
	var vals [3]int
	vals[0] = 12
	vals[1] = 1
	fmt.Println(vals)
	// declare and assign
	more := [2]int{1, 4}
	fmt.Println(more)
	// `fruits` is a slice. slice is like C++ vector, dynamic size
	fruits := []string{"Apple", "Komola", "Kola", "AngularJS"}
	fmt.Println(len(fruits))
	fmt.Println(fruits[0])
	fmt.Println(fruits[1:3]) // form: [)
	// so arrays are of the form [n]T where n is its size
	// slices are of the form []T, since size is dynamic
	// seems like python np array
	fmt.Println(fruits[0:])
	fmt.Println(fruits[:])
	fmt.Println(fruits[:4])
}
