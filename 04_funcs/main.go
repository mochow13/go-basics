package main

import "fmt"

func sayName(name string) string {
	return "That's mah city " + name
}

func getSum(a int32, b int32) int32 {
	return a + b
}

// can avoid writing int32 twice because a and b are same type
func getMultiplication(a, b int32) int64 {
	return int64(a) * int64(b) // typecast
}

// any number of return values in go
func divByTwo(a int) (int, int) {
	return a / 2, a % 2
}

// naked return
// means that we return two variables x and y of type int
func division(a, b int) (x, y int) {
	x = a / b
	y = a % b
	return // this return functions returns the named return params
}

func main() {
	fmt.Println(sayName("chittagong"))
	fmt.Println(getSum(32, 552))
	fmt.Println(getMultiplication(100000, 100000))
	fmt.Println(divByTwo(3421))
	fmt.Println(division(442, 14))
}
