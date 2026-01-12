package main

import (
	"fmt"
)

var x int

func main() {

	fmt.Println(x)

	y := 42

	fmt.Println(y)
	_, z := -1, -2

	fmt.Println(z)

	var p float32 = 3.14
	fmt.Printf("%v is of type: %T \n", p, p)

	fmt.Println(`
	repetition 
	is the
	mother of mastery`)

}

/*
var for zero value
short declaration operator
multiple initializations
var when specificity is required
blank identifier
*/
