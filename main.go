package main

import (
	"fmt"
	"sort"

	"training.go/sort/input"
)

func main() {

	fmt.Println("--- This program allows to sort a list of values provided by the user ---")
	fmt.Println("--- It will then indicate the min and max values ---")
	values := input.ReadInputs()
	fmt.Printf("--- Provided values are: %v ---\n", values)
	sort.Ints(values)
	fmt.Printf("--- Sorted values are: %v ---\n", values)
	fmt.Printf("--- Min value is: %v ---\n", values[0])
	fmt.Printf("--- Max value is: %v ---\n", values[4])
	fmt.Println("--- End of program ---")
}
