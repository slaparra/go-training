package main

import "fmt"

var z int = 3

func main() {
	x := "some string value"
	printString(x)

	y := getClosure()
	fmt.Printf("Z value: %d\n", y())
	fmt.Printf("Z value: %d\n", y())
}

func printString(s string) {
	fmt.Printf("Print string variable: %s\n", s)
}

func getClosure() func() int {
	return func() int {
		z++
		return z
	}
}

//Result:
//	Print string variable: some string value
//	Z value: 4
//	Z value: 5
