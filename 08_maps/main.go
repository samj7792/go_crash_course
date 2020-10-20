package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign keys and values
	emails["bob"] = "bob@gmail.com"
	emails["will"] = "will@gmail.com"
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["bob"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "bob")
	fmt.Println(emails)

	// Declare map and add key values
	nicknames := map[string]string{"william": "will", "Robert": "Rob"}
	nicknames["Nicholas"] = "nick"
	fmt.Println(nicknames)
}
