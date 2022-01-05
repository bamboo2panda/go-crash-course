package main

import "fmt"

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s.", a.Name, a.Sound)
	fmt.Println()
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs.", a.Name, a.NumberOfLegs)
	fmt.Println()
}

func main() {
	z := addTwoNumbers(2, 4)
	fmt.Println(z)

	myTotal := sumMany(2, 3, 4, 5, 6)
	fmt.Println(myTotal)

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "Woof"
	dog.NumberOfLegs = 4
	dog.Says()
	dog.HowManyLegs()

	cat := Animal{
		Name:         "cat",
		Sound:        "Myau",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManyLegs()
}

func addTwoNumbers(x, y int) int {
	return x + y
}

func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}
	return total
}
