package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %v: %v!", g.LanguageName(), g.Greet(name))
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(Name string) string {
	return fmt.Sprintf("Ciao %v", Name)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(Name string) string {
	return fmt.Sprintf("Olá %v", Name)
}

type English struct {
}

func (e English) LanguageName() string {
	return "English"
}

func (e English) Greet(Name string) string {
	return fmt.Sprintf("Howdy %v", Name)
}
