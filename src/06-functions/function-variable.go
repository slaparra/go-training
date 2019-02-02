package main

import (
	"fmt"
)

var g func()

// x := 7 outside main can not assign values to a variable without declare it, only declare variables

var x int = 7

func main() {
	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("f type: %T\n", f)
	fmt.Printf("g type: %T\n", g)

	g := f
	g()

	fmt.Printf("x type: %T\n", x)
}
