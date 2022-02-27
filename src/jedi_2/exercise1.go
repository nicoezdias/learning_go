/*
Hands-on exercise #1
Write a program that prints a number in decimal, binary, and hex
*/
package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Printf("Decimal: %d\nBinary: %b\nHex: %#x\n", a, a, a)
}
