// --summary:
// use functions to perform basic operations and print some informtaion to the terminal

// --requirements:
// *Write a function that accepts a person's name as a function parameter and displayas a greeeting to that person.
// *Write a function that returns any message, and call it from within fmt.PrintLn()
// *Write a function to add 3 numbers togather, supplied as a argument ,and return the answer
// *Write a function that returns any number
// *Write a function that returns any two numbers
// *Add three numbers togather using any combination of existing functions .Print the result
// *Call every function at least once
package main

import "fmt"

// *Write a function that accepts a person's name as a function parameter and displayas a greeeting to that person.
func greet(name string) {
	fmt.Printf("Hello %v. Greetings!\n", name)
}

// *Write a function that returns any message, and call it from within fmt.PrintLn()
func message() string {
	return "Hi there. This is a message"
}

// *Write a function to add 3 numbers togather, supplied as a argument ,and return the answer
func add3numbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

// *Write a function that returns any number
func returnanynumber() int {
	return 1
}

// *Write a function that returns any two numbers
func returnanytwonumber() (int, int) {
	return 2, 2
}
func main() {
	fmt.Println("Starting functions excercise...")
	greet("Alice")
	fmt.Println(message())
	a, b := returnanytwonumber()
	answer := add3numbers(returnanynumber(), a, b)
	fmt.Printf("Result: %d\n", answer)
}
