package main

import (
	"fmt"
	"strings"
)

func main() {
	newString := "Go is a great programming language. Go for it!"

	if strings.Contains(newString, "Go") {
		// newString = strings.Replace(newString, "Go", "Golang", 1)
		newString = strings.ReplaceAll(newString, "Go", "Golang")
	}

	fmt.Println(newString)

	// string comparison
	if "Alpha" > "Absolute" {
		fmt.Println("A is greater then B")
	} else {
		fmt.Println("A is not greater than B")
	}

	badEmail := " me@here.com "
	badEmail = strings.TrimSpace(badEmail)
	fmt.Printf("=%s=", badEmail)
	fmt.Println()

}
