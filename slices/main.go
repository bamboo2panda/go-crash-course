package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "fish")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Element 0 is", animals[0])
	fmt.Println("Firt two elements are", animals[0:2])
	fmt.Println("The slice is", len(animals), "elements long.")
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	DeleteFromSlice(&animals, 1)
	fmt.Println(animals)

}

func DeleteFromSlice(a *[]string, i int) {
	ar := *a
	ar[i] = ar[len(ar)-1]
	ar[len(ar)-1] = ""
	*a = ar[:len(ar)-1]
}
