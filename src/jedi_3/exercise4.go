/*
Hands-on exercise #4
Create a for loop using this syntax
	- for{}
Have it print out the years you have been alive.
*/

package main

import (
	"fmt"
)

func main() {
	alive := 1999
	for {
		if alive > 2022 {
			break
		}
		fmt.Println(alive)
		alive++
	}
}
