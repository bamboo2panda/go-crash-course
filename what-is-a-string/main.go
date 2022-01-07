package main

import "fmt"

func main() {
	courseName := "Learn Go for Beginners Crash course"

	for i := 0; i <= 21; i++ {
		fmt.Println(string(courseName[i]))
	}
	fmt.Println()

	fmt.Println("Length of courseName is", len(courseName))

	var mySlice []string
	mySlice = append(mySlice, "one")
	mySlice = append(mySlice, "two")
	mySlice = append(mySlice, "three")

	fmt.Println("mySlice has", len(mySlice), "elements")
	fmt.Println("the last element in mySlice is", mySlice[len(mySlice)-1])
}
