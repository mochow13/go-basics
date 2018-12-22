package main

import "fmt"

func main() {
	info := make(map[string]int) // key value
	info["mochow"] = 1993
	info["sam"] = 1944
	info["zfr"] = 1994
	fmt.Println(info)

	// initialize while declaring
	data := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(data)

	delete(info, "zfr")
	fmt.Println(info)
}
