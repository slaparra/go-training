package main

import "fmt"

/*
	Write a function which takes an integer. The function will have two returns.
	- The first return should be the argument divided by 2.
	- The second return should be a bool that let’s us know whether or not the argument was even. For example:
		half(1) returns (0, false)
		half(2) returns (1, true)
*/
func main() {
	var input int

	fmt.Printf("Enter a number: ")

	fmt.Scan(&input)
	aVar, anotherVar := twoParameterReturnedFunction(input)

	var evenOrOdd string

	if anotherVar {
		evenOrOdd = "even"
	} else {
		evenOrOdd = "odd"
	}

	fmt.Printf("%d divided by 2 is %d and this number is %s\n", input, aVar, evenOrOdd)
}

func twoParameterReturnedFunction(input int) (int, bool) {
	return input / 2, input%2 == 0
}
