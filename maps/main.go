package main

import "fmt"

func main() {
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, val := range intMap {
		fmt.Println(key, val)
	}

	// delete(intMap, "four")
	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")
	}

	intMap["two"] = 4

	for key, val := range intMap {
		fmt.Println(key, val)
	}
}