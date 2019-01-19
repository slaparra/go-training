package main

import (
	"errors"
	"fmt"
)

var ErrMath = errors.New("math: square root of negative number") //idiomatic way to declare an error

func main() {
	fmt.Printf("%T\n", ErrMath) //*errors.errorString

	_, err := Sqrt(-10.98)
	if err != nil {
		fmt.Println(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrMath
		//return 0, fmt.Errorf("Show the error with format: %v\n", f)
	}

	return 8.3, nil
}
