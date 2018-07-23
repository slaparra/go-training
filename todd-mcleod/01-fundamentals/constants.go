package main

import "fmt"

const aConstantString string = "a string constant value"
const (
	Sunday = iota //https://golang.org/ref/spec#Iota , iota only for constants
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays //this const is not exported
)

func main() {
	const aConstantInt int = iota

	fmt.Printf("Printing a constant string: %s\n", aConstantString)
	fmt.Printf("Printing a constant integer: %d\n", aConstantInt)

	fmt.Printf("Print monday %v\n", Monday)
	fmt.Printf("Print saturday %v\n", Saturday)
	fmt.Printf("Print number of days %v\n", numberOfDays)
	fmt.Printf("iota type %T\n", Sunday)

	//aConstantInt = 4 //cannot assign a value to a constant
}
