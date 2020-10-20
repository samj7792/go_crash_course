package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	vegArray := [2]string{"carrot", "broccoli"}

	fruitSlice := []string{"apple", "orange", "grape", "banana"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(vegArray)
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
