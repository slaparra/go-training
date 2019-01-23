package main

import "fmt"

func main() {
	var a int
	fmt.Print("sum 1: ")
	fmt.Scan(&a)
	var b int
	fmt.Print("sum 2: ")
	fmt.Scan(&b)

	fmt.Printf("Result %d + %d = %d\n", a, b, Sum(a, b))
}
