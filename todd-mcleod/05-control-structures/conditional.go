package main

import "fmt"

func main() {
	x := askForAnInteger()

	if x % 2 == 1 {
		fmt.Printf("X (%d) is odd\n", x)
	} else {
		fmt.Printf("X (%d) is even\n", x)
	}

	//with initialization statement
	//in this case, b is only inside the scope of the if conditional

	if b := askForAnInteger(); b % 2 == 1 {
		fmt.Printf("X (%d) is odd\n", b)
	} else {
		fmt.Printf("X (%d) is even\n", b)
	}

	//here we can't access to the variable 'b'
}

func askForAnInteger() int {
	fmt.Print("Write an integer: ")
	var x int
	fmt.Scan(&x)
	return x
}
