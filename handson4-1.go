package main

import "fmt"

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
	fmt.Printf("Decimal: %d \t Binary: %b\n", 1, 1)
	fmt.Printf("Decimal: %d \t Binary: %b \n", 1<<1, 1<<1)
	fmt.Printf("Decimal: %d \t Binary: %b \n", 1<<2, 1<<2)
	fmt.Printf("Decimal: %d \t Binary: %b \n", 1<<3, 1<<3)
	fmt.Printf("Decimal: %d \t Binary: %b \n", 1<<4, 1<<4)
	fmt.Printf("Decimal: %d \t Binary: %b \n", 1<<5, 1<<5)
	fmt.Printf("Decimal: %d \t Binary: %b \n", 1<<6, 1<<6)
	fmt.Printf("Decimal: %d \t Binary: %b \n", 1<<7, 1<<7)
	fmt.Printf("Decimal: %d \t Binary: %b \n", 1<<8, 1<<8)
}
