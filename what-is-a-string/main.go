package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	name := "Hello world"
	fmt.Println("String:", name)
	fmt.Println()
	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}
	fmt.Println()

	name = "Γειά σου Κόσμε"
	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}
	fmt.Println()

	fmt.Println()
	fmt.Println("Tree ways to concatinate strings")

	h := "Hello "
	w := "World."

	// using + not very efficient
	mystring := h + w
	fmt.Println(mystring)
	fmt.Println()

	// using fmt - more efficient
	mystring = fmt.Sprintf("%s%s", h, w)
	fmt.Println(mystring)
	fmt.Println()

	// using stringbuilder - very efficient
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())

	fmt.Println()
	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println("Getting substring")
	fmt.Println(name[1:13])
}
