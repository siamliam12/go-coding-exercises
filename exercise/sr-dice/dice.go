// --Summary:
// Create a program that can perform dice rolls. The program should report the results as detailed in the requirements

// --Requirements:
// *Print the sum of the dice roll
// *Print additional information in these circumstances:
// -"snake eyes": when the total roll is 2, and total dice is 2
// -"lucky 7":when the total roll is 7
// -"Even":when the total roll is even
// -"Odd":when the total roll is odd
// *The program must hanbdle any number of dice,rolls, and sides that

// --Notes:
// *Use packeges form the tsandard library to complete the program
package main

import (
	"fmt"
	"math/rand"
	"time"

)
func roll(sides int) int {
	return rand.Intn(sides) + 1
}
func main() {
	rand.Seed(time.Now().UnixNano())
	dice, sides := 2,12
	rolls := 1
	for r:= 1,r<= rolls; r++ {
		sum := 0
		for d:=1;d<= dice;d++ {
			rolled := roll(sides) 
			sum += rolled
			fmt.PrintLn("roll #", r,"Die #",d,":",rolled)
		}
		fmt.PrintLn("Total rolled:",sum)
		switch  sum := sum {
		case sum ==2 && dice == 2:
			fmt.PrintLn("Snake eyes")
		case sum == 7:
			fmt.PrintLn("Lucky 7")
		case sum %2 == 0:
			fmt.PrintLn("Even")
		case sum %2 == 1:
			fmt.PrintLn("Odd")
		}
	}
}