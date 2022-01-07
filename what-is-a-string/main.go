package main

import (
	"fmt"
	"strings"
)

func main() {
	courses := []string{
		"Learn Go for beginners Crash cource",
		"Learn Java for beginners Crash cource",
		"Learn Python for beginners Crash cource",
		"Learn C for beginners Crash cource",
	}

	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found on", x, "and index is", strings.Index(x, "Go"))
		}
	}

	newString := "Go is a great programming language. Go for it!"
	fmt.Println(strings.HasPrefix(newString, "Go"))
	fmt.Println(strings.HasPrefix(newString, "Python"))
	fmt.Println(strings.HasSuffix(newString, "!"))
	fmt.Println(strings.Count(newString, "Go"))
	fmt.Println(strings.Count(newString, "Fish"))
	fmt.Println(strings.Index(newString, "Go"))
	fmt.Println(strings.Index(newString, "Python"))
	fmt.Println(strings.LastIndex(newString, "Go"))
}
