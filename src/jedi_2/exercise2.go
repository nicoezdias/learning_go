/*
Hands-on exercise #2
Using the following operators, write expressions and assign their values to variables:
	g. ==
	h. <=
	i. >=
	j. !=
	k. <
	l. >
Now print each of the variables.
*/
package main

import (
	"fmt"
)

func main() {
	g := 70 == 70
	h := 30 <= 48
	i := 48 >= 48
	j := 70 != 70
	k := 56 < 55
	l := 56 > 55

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}
