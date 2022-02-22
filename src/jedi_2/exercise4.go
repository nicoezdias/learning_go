package main

import (
	"fmt"
)

func main() {
	a := 64
	fmt.Printf("Decimal: %d\t\tBinary: %b\t\tHex: %#x\n", a, a, a)
	b := a << 1
	fmt.Printf("Decimal: %d\t\tBinary: %b\tHex: %#x\n", b, b, b)
}
