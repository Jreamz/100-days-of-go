package main

import "fmt"

const (
	repetition string = "repetition"
)

var mastery string = "mastery"

func main() {
	sentence := repetition + " is the mother of " + mastery
	fmt.Println(sentence)
}
