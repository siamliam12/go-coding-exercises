//summary:
// print basic information to the terminal using various variable creating teachniques. The information may be printed using any formatting you like.

// requirements:
// Store your favorite color in a variable using the `var` keyword
// Store your birth year and age (in year) in two variable using compound assignment
// Store your first and last initials in two variables using block assignment
// declare (but don't assign ) a variable for your age (in days)
// then assign it on the next line by multiplying 365 with the age variable created earlier

// notes:
// Use fmt.Println() to print out information
// basic math operations are
// subtraction -
// addition +
// multiplication *
// division /

package main

import "fmt"

func main() {
	fmt.Println("Starting variable excercise...")
	var favoriteColor = "Black"
	fmt.Println("favoriteColor: ", favoriteColor)
	birthYear, age := 2003, 20
	fmt.Println("birthYear: ", birthYear, " age=", age)
	var (
		firstInitial = "S"
		lastInitial  = "A"
	)
	fmt.Println("first: ", firstInitial, " last: ", lastInitial)
	var myAge int
	myAge = age * 365
	fmt.Println("myAge: ", myAge)
}
