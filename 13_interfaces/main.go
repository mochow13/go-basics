package main

import (
	"fmt"
	"math"
)

// Interfaces are named collection of method signatures
// It's also a type
type shapeOfYou interface {
	perim() float64
}

// any struct that implements perim() will actually use the interface
type circle struct {
	radius float64
}

type recta struct {
	w, h float64
}

// implementing perim() that returns float64 val
func (r recta) perim() float64 {
	return 2*r.w + 2*r.h
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func basic(s shapeOfYou) float64 {
	return s.perim()
}

func main() {
	Circle := circle{12}
	Rectang := recta{5, 14}
	// fmt.Println(Circle.perim())
	// fmt.Println(Rectang.perim())
	fmt.Println(basic(Circle))
	fmt.Println(basic(Rectang))
}
