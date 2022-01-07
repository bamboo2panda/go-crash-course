package main

import "fmt"

func main() {
	courseName := "Learn Go for Beginners Crash course"
	fmt.Println(string(courseName[0]))
	fmt.Println(string(courseName[6]))

	for i := 0; i <= 21; i++ {
		fmt.Println(string(courseName[i]))
	}
	fmt.Println()

	for i := 13; i <= 21; i++ {
		fmt.Println(string(courseName[i]))
	}
	fmt.Println()
}
