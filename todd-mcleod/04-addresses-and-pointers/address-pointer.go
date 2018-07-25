package main

import "fmt"

func main() {
	var x int = 1
	var p *int //pointer to an int

	p = &x //assign the address of x to the pointer p

	//here the content of pointer p is equals to the value of x
	fmt.Printf("the pointer p (%v) is equals to x address (%v)\n", p, &x)
	fmt.Printf("and the content of pointer p (%v) is equals to the x value (%v)\n", *p, x)

	*p = 2
	fmt.Printf("Type pointer: %T\n", p)
	fmt.Printf("so if we change the content of the pointer p (%v) we are changing the value of x (%v)", *p, x)
}
