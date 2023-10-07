package main

import "fmt"

func average(a, b, c int) float32 {
	return float32(a+b+c) / 3
}

func main() {
	quize1, quize2, quize3 := 9, 7, 8
	if quize1 > quize2 {
		fmt.Println("quize1 scored higher than quiz2")
	} else if quize1 < quize2 {
		fmt.Println("quize2 scored higher than quiz1")
	} else {
		fmt.Println("quize1 and quiz2 have the same score")
	}

	if average(quize1, quize2, quize3) > 7 {
		fmt.Println("acceptable grades")
	} else {
		fmt.Println("Instructor did a bad job!")
	}
}
