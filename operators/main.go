package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 12
	b := 6

	c, err := devideTwoNumbers(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		if c == 2 {
			fmt.Println("We found 2!")
		}
	}
}

func devideTwoNumbers(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("can not devide by zero")
	}
	return x / y, nil
}
