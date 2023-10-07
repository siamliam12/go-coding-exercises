package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Printf("Hello from greet function\n")
}

func main() {
	greet()
	dozen := double(6)
	fmt.Printf("A dozen is %v\n", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Printf("A baker's dozen is %v\n", bakersDozen)
	anotherBakersDozen := add(double(6), 1)
	fmt.Printf("Another baker's dozen is %v\n", anotherBakersDozen)
}
