package main

import (
	"fmt"
	"strconv" // string converter
)

// Person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
// We need to convert the age to a string with strconv.Itoa
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	}
	p.lastName = spouseLastName

}

func main() {
	// Initialize person using struct
	person1 := Person{firstName: "Sam", lastName: "Jackson", city: "Fairfax", gender: "m", age: 28}

	// Alternatively
	person2 := Person{"Sara", "Smith", "Fairfax", "f", 28}

	fmt.Println(person1)
	fmt.Println(person2.firstName)

	person2.age++
	fmt.Println(person2)

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person2.getMarried("Williams")
	fmt.Println(person2.greet())
}
