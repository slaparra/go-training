package main

import "fmt"

// from: https://blog.golang.org/defer-panic-and-recover

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	//if a method can result in a panic:
	//  put a defer inside a function
	//  and call the recover method to don't break the execution
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
