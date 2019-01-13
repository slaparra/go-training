package main

import (
	"fmt"
)

// https://www.udemy.com/learn-how-to-code/learn/v4/overview
// file from https://github.com/GoesToEleven/GolangTraining
// goo.gl/flGsyX

/*
func main() {
	f := factorial(4)
	fmt.Println("Total:", f)
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
*/

/*
CHALLENGE #1:
-- Use goroutines and channels to calculate factorial

CHALLENGE #2:
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your own answer, then post that answer to this discussion area: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reasons of others
*/

func main() {
	f := factorial(60)
	fmt.Println("Total:", f)
}

func factorial(n int) uint64 {
	channel := reduceNumber(n)

	return mult(channel)
}

func reduceNumber(n int) chan int {
	c := make(chan int)
	go func() {
		for i := n; i > 0; i-- {
			c <- i
		}
		close(c)
	}()

	return c
}

func mult(factorialNumbers chan int) uint64 {
	var result uint64 = 1
	for numbers := range factorialNumbers {
		result *= uint64(numbers)
	}

	return result
}
