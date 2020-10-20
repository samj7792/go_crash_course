package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint 16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Sam"
	var age = 28
	var isCool = true
	isCool = false // vars can be redefined, but const cannot

	// Shorthand
	// name := "Sam"
	// email := "sam@gmail.com"

	name, email := "Sam", "sam@gmail.com"

	fmt.Println(name, email, age, isCool)
	fmt.Printf("%T\n", name)
}
