package thefarm

import "fmt"

type InvalidCowsError struct {
	errorMessage string
	cows         int
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.errorMessage)
}

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, cows int) (float64, error) {
	fatteningFactor, err := f.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	fodderAmount, err := f.FodderAmount(cows)
	if err != nil {
		return 0.0, err
	}

	return fatteningFactor * fodderAmount / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(f, cows)
	} else {
		return 0.0, fmt.Errorf("invalid number of cows")
	}

}

func ValidateNumberOfCows(cows int) error {
	if cows > 0 {
		return nil
	} else if cows < 0 {
		return &InvalidCowsError{cows: cows, errorMessage: "there are no negative cows"}
	} else {
		return &InvalidCowsError{cows: cows, errorMessage: "no cows don't need food"}
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
