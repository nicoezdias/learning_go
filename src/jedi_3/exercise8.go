/*
Hands-on exercise #8
Create a program that uses a switch statement with no switch expression specified.
*/

package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("not print")
	case true:
		fmt.Println("print")
	}
}
