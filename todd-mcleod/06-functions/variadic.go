package main

import "fmt"

func main() {

	result := sumSlice(1.5, 2.6, 3.7) //some arguments
	fmt.Printf("%.2f\n", result)

	aSlice := []float64{1.5, 2.6, 3.7}

	result = sumSlice(aSlice...) //slice to variadic
	fmt.Printf("%.2f\n", result)
}

func sumSlice(someFloats ...float64) (result float64) { //receive arguments
	fmt.Println(someFloats)
	fmt.Printf("type of someFloats: %T\n", someFloats) // []int

	for index, value := range someFloats {
		result += value
		fmt.Printf("Index: %d - value: %.2f\n", index, value)
	}

	return
}
