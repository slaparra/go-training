package main

import "fmt"

func main() {
	x := askForAnInteger()

	if x%2 == 1 {
		fmt.Printf("X (%d) is odd\n", x)
	} else {
		fmt.Printf("X (%d) is even\n", x)
	}
}

func askForAnInteger() int {
	fmt.Print("Write an integer: ")
	var x int
	fmt.Scan(&x)
	return x
}
