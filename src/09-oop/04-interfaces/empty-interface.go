package main

import "fmt"

func main() {
	var a interface{}

	a = 2.3
	printA(a)

	a = 33
	printA(a)

	a = "James"
	printA(a)

	type dog struct {
		age  int
		name string
	}

	a = dog{3, "Bobby"}
	printA(a)
}

func printA(a interface{}) {
	fmt.Printf("Variable a, with value %v is type %T\n", a, a)
}
