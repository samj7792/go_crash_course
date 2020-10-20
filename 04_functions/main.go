package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(x int, y int) int { // args can be shortened to (x, y int)
	return x + y
}

func main() {
	fmt.Println(greeting("Sam"))
	fmt.Println(getSum(1, 2))
}
