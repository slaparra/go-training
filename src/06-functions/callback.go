package main

import (
	"fmt"
)

var g func()

func main() {
	f := func(xi []int) int {
		var result int
		for _, i := range xi {
			result += i
		}
		return result
	}

	aResult := foo([]int{1, 3, 4}, f)

	fmt.Println(aResult)

}

func foo(slice []int, callback func(xi []int) int) int {
	return callback(slice)
}
