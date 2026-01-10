package main

import "fmt"

const (
	_ = 1 << iota
	b1
	b2
	b3
	b4
	b5
	b6
	b7
	b8
)

func main() {
	/*
		Bitwise operators, we are shifting the bits left (or you can do right)
		This means that 1 << 1 is 2
		1 << 2 is 4
		1 << 3 is 8
		1 << 4 is 16
		1 << 5 is 32
		1 << 6 is 64
		1 << 7 is 128
		1 << 8 is 256
		etc
	*/
	fmt.Printf("Decimal: %d \t Binary: %b \n", b1, b1)
	fmt.Printf("Decimal: %d \t Binary: %b \n", b2, b2)
	fmt.Printf("Decimal: %d \t Binary: %b \n", b3, b3)
	fmt.Printf("Decimal: %d \t Binary: %b \n", b4, b4)
	fmt.Printf("Decimal: %d \t Binary: %b \n", b5, b5)
	fmt.Printf("Decimal: %d \t Binary: %b \n", b6, b6)
	fmt.Printf("Decimal: %d \t Binary: %b \n", b7, b7)
	fmt.Printf("Decimal: %d \t Binary: %b \n", b8, b8)
}
