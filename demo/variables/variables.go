package main

import (
	"fmt"
)

func main() {
	var myName = "siam"
	fmt.Println("my name is ", myName)
	var name string = "Kathy"
	fmt.Println("name=", name)

	userName := "admin"
	fmt.Println("username =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("part1 =", part1, "other is", other)

	part2, other := 2, 0
	fmt.Println("part2 =", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("The sum is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName =", lessonName)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}
