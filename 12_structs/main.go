package main

import "fmt"

type info struct {
	username string
	pass     string
}

type rectangle struct {
	height int
	width  int
}

// struct method value receiver
func (r rectangle) area() int {
	return r.width * r.height
}

// struct method pointer receiver

func (r *rectangle) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	user1 := info{"abc", "123"}
	fmt.Println(user1)
	user2 := info{username: "bcd", pass: "1234"}
	fmt.Println(user2)
	user2.pass = "wow"
	fmt.Println(user2) // changes
	var ptr *info
	ptr = &user2
	fmt.Println(*ptr)
	ptr.pass = "wwowowow"
	fmt.Println(user2)

	rect := rectangle{4, 6}
	fmt.Println(rect.area())
	fmt.Println(rect.perim())
}
