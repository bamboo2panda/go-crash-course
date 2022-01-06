package main

import (
	"fmt"
	"math"
)

func main() {
	//multiplication
	// area = πr
	var radius = 12.0
	area := math.Pi * radius * radius
	fmt.Println("area is", area)

	// integer division
	half := 1 / 2
	fmt.Println("half with integer division:", half)

	halfFloat := 1.0 / 2.0
	fmt.Println("halfFloat:", halfFloat)

	// squering (raising to the power)
	badThreeSquared := 3 ^ 2
	fmt.Println("Bad three squared", badThreeSquared)

	goodThreeSquared := math.Pow(3.0, 2.0)
	fmt.Println("Good three squared", goodThreeSquared)

	// modulus
	remainder := 50 % 3
	fmt.Println("Remainder:", remainder)

	// unary operators
	x := 3
	x++
	fmt.Println("x is now", x)

	x--
	x--
	fmt.Println("x is now", x)
}
