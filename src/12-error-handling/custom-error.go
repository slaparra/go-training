package main

import "fmt"

// ---------------------

//Custom type error

// Naming convention for custom error types:
// --> It is idiomatic in Go to postfix the name of a custom error type with the word Error.
type CustomError struct {
	param    int
	errorMsg string
	method   string
}

// implement method Error() to follow error interface: https://golang.org/pkg/errors/
func (error *CustomError) Error() string {
	return fmt.Sprintf("Error: %d %s - Method: %s", error.param, error.errorMsg, error.method)
}

// ---------------------

func main() {
	var scannedVar int
	fmt.Print("Find X -> 20/X + 10/5 = 6: ")
	fmt.Scan(&scannedVar)

	err := checkSolution(scannedVar)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("20/%d + 10/5 = 6 -> %d is correct\n", scannedVar, scannedVar)
	}
}

func checkSolution(scannedVar int) error {
	if scannedVar == 0 {
		panic("(Panic msg): X must not be zero")
	}

	// returns the custom error
	if scannedVar != 5 {
		return &CustomError{scannedVar, "is not the expected value", "checkSolution"}
	}

	return nil
}
