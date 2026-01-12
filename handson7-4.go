package main

import "fmt"

func main() {
	var x int8 = 127
	var y uint8 = 255
	fmt.Printf("%v is of type %T \n", x, x)
	fmt.Printf("%v is of type %T \n", y, y)
}

/*
declare two vars
one var is of VALUE int8
second var is of TYPE uint8
assign each to their largest possible values then print

how I found these values
https://go.dev/ref/spec#Numeric_types
*/
