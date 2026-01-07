package main

import (
	"fmt"
)

func main() {
	//adams := 42
	////fmt.Printf("42 as binary, %b \n", adams)
	////fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("%v value is \t binary: %b \t hexidecimal: %#x \n", a, a, a)
	fmt.Printf("%v value is \t binary: %b \t hexidecimal: %#x \n", b, b, b)
	fmt.Printf("%v value is \t binary: %b \t hexidecimal: %#x \n", c, c, c)
	fmt.Printf("%v value is \t binary: %b \t hexidecimal: %#x \n", d, d, d)
	fmt.Printf("%v value is \t binary: %b \t hexidecimal: %#x \n", e, e, e)
	fmt.Printf("%v value is \t binary: %b \t hexidecimal: %#x \n", f, f, f)
}
