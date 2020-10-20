package main

import "fmt"

func main() {
	// create slice
	ids := []int{33, 25, 76, 34, 83}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids { // if you're not printing i you'll get an error, so use _ instead
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// range with map
	emails := map[string]string{"bob": "bob@gmail.com", "mike": "mike@gmail.com", "will": "will@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
