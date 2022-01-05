package main

import (
	"fmt"
	"log"
)

var myInt int

var myInt16 int16
var myInt32 int32
var myInt64 int64

var myFloat float32
var myFloat64 float64

var myUint uint //0, 1, 2, 3 ...

var myStrings [3]string

type Car struct {
	NumberOfTypes int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	myInt = 10
	myUint = 20
	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Trevor"

	log.Println(myString)

	myString = "John"

	var myBool = true
	myBool = false

	log.Println(myBool)

	// myString[0] = "Cat"
	// myString[1] = "Dog"
	// myString[2] = "Fish"

	var myCar Car

	myCar.NumberOfTypes = 4
	myCar.Luxury = false
	myCar.Make = "Volkswagen"

	nextCar := Car{
		NumberOfTypes: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}

	fmt.Printf("My car is %d %s %s", nextCar.Year, myCar.Make, myCar.Model)

}
