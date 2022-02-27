/*
Hands-on exercise #7
Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
*/

package main

import (
	"fmt"
)

func main() {
	a := 42
	b := 35
	c := 0

	if a == 0 {
		fmt.Printf("%v is zero\n", a)
	} else if a%2 == 0 {
		fmt.Printf("%v is even\n", a)
	} else {
		fmt.Printf("%v is odd\n", a)
	}

	if b == 0 {
		fmt.Printf("%v is zero\n", b)
	} else if b%2 == 0 {
		fmt.Printf("%v is even\n", b)
	} else {
		fmt.Printf("%v is odd\n", b)
	}

	if c == 0 {
		fmt.Printf("%v is zero\n", c)
	} else if c%2 == 0 {
		fmt.Printf("%v is even\n", c)
	} else {
		fmt.Printf("%v is odd\n", c)
	}
}
