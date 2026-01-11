package main

import "fmt"

// ByteSize Create type for byte size
type ByteSize int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	// Print out the binary values of the constants
	fmt.Printf("%d KB = %b \n", 1, KB)
	fmt.Printf("%d MB = %b \n", 1, MB)
	fmt.Printf("%d GB = %b \n", 1, GB)
	fmt.Printf("%d TB = %b \n", 1, TB)
	fmt.Printf("%d PB = %b \n", 1, PB)
	fmt.Printf("%d EB = %b \n", 1, EB)
}
