package main

import "fmt"

func modify(x int) { // param passed by value
	x = x * 2
}

func modifyByRef(x *int) { // still passed by value, pointer is copied
	*x = *x * 2
}

// Go does not pass anything by reference. When we pass a pointer
// it copies it. Since it works on memory addresses, changes are reflected accordingly.
// Go has no pointer arithmetic, unlike C!
func main() {
	var a int
	var ptr *int
	a = 13
	ptr = &a
	fmt.Println(ptr, a)
	*ptr = 26
	fmt.Println(ptr, a)
	fmt.Printf("%T %T\n", ptr, a)
	a = 34
	fmt.Println(*ptr)
	modify(a)
	fmt.Println(*ptr, a)
	modifyByRef(ptr)
	fmt.Println(*ptr, a)
}
