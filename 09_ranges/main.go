package main

import "fmt"

func main() {
	// ranges are like iterators in C++
	// iterate through various data structures in Go
	vals := []int{2, 4, 4, 1, 1, 3, 9}
	sum := 0
	for i, val := range vals {
		fmt.Println(i, val) // i is index
		sum += val * (i + 1)
	}
	fmt.Println(sum)
	// what if we don't care about indices
	sum = 0 // not :=
	for _, val := range vals {
		sum += val
	}
	fmt.Println(sum)

	// iterate on maps
	info := map[string]int{"foo": 1, "bar": 2, "buzz": 3}
	for key, value := range info {
		fmt.Printf("%s -> %d\n", key, value)
	}

	for i, ch := range "randomstring" {
		fmt.Println(i, ch) // index and rune
	}

	// want only the keys?
	for key := range info {
		fmt.Printf("key: %s\n", key)
	}

	// only the values?
	for _, val := range info {
		fmt.Printf("value: %d\n", val)
	}
}
