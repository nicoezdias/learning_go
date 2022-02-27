/*
Hands-on exercise #3
Create a for loop using this syntax
	- for condition { }
Have it print out the years you have been alive.
*/

package main

import (
	"fmt"
)

func main() {
	alive := 1999
	for alive <= 2022 {
		fmt.Println(alive)
		alive++
	}
}
