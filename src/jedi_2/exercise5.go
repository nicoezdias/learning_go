/*
Hands-on exercise #5
Create a variable of type string using a raw string literal. Print it.
*/

package main

import (
	"fmt"
)

func main() {
	a := `There are only 10 types 
of people in this world;
those who understand binary 
and those who don't.`

	fmt.Println(a)
}
