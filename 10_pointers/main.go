package main

import "fmt"

func main() {
	a := 5
	b := &a // assign b to a pointer of a or its memory address

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(*b)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
