/*
Hands-on exercise #6
Create a program that shows the “if statement” in action.
*/

package main

import (
	"fmt"
)

func main() {
	a := 42
	b := 35
	if a%2 == 0 {
		fmt.Printf("%v is even\n", a)
	}
	if b%2 != 0 {
		fmt.Printf("%v is odd\n", b)
	}
}
