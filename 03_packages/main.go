package main

import (
	"fmt"
	"math"

	"github.com/mochow/basics/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(3.2))
	fmt.Println(math.Sqrt(144))
	fmt.Println(math.MaxInt64)
	fmt.Println(string(0x2318))
	fmt.Println(strutil.Reverse("mochow"))
}
