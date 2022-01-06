package main

import (
	"fmt"
	"myapp/packageone"
)

var one = "The package variable"

func main() {

	var somefingElse = "The block variable"
	fmt.Println(somefingElse)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	packageone.Exported()
}

func myFunc() {

	fmt.Println(one)
}
