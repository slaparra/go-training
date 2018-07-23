package main

import "fmt"

func main() {
	//Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).
	var x int = 1
	var y int = 2

	fmt.Printf("X address is %v and value %d\n", &x, x)
	fmt.Printf("Y address is %v and value %d\n", &y, y)

	swap(&x, &y)

	fmt.Printf("X after swap is %d\n", x)
	fmt.Printf("Y after swap is %d\n", y)
}

func swap(xAddress *int, yAddress *int) {
	fmt.Printf("address of xAddress before swap: %v\n", xAddress)
	fmt.Printf("address of yAddress before swap: %v\n", yAddress)
	tmp := *xAddress
	fmt.Printf("tmp is xAddress: %v\n", tmp)
	*xAddress = *yAddress
	*yAddress = tmp
}
