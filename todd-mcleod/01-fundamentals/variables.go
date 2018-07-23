package main

import "fmt"

var anInteger int = 10
var aString string = "a string value"
var aFloat float32

func main() {
	aBoolean := true
	anotherFloat := 4.5

	fmt.Printf("%v\n", anInteger) // %v the value in a default format
	fmt.Printf("%T\n", anInteger) // %T a Go-syntax representation of the type of the value
	fmt.Printf("%v\n", aBoolean)
	fmt.Printf("%T\n", aBoolean)
	fmt.Printf("%v\n", aString)
	fmt.Printf("%T\n", aString)
	fmt.Printf("%v\n", aFloat) // aFloat has the default value for floats because is not initialized
	fmt.Printf("%T\n", aFloat)
	fmt.Printf("%v\n", anotherFloat)
	fmt.Printf("%T\n", anotherFloat)
	fmt.Printf("%10.2f\n", 520702.23723525)

	//without initialize variables
	var a2 string
	var b2 float32
	var c2 bool
	var d2 int

	fmt.Printf("print default values for:\n")
	fmt.Printf("string %v\n", a2)
	fmt.Printf("float %v\n", b2)
	fmt.Printf("bool %v\n", c2)
	fmt.Printf("int %v\n", d2)

	/*
		bool:                    %t
		int, int8 etc.:          %d
		uint, uint8 etc.:        %d, %#x if printed with %#v
		float32, complex64, etc: %g
		string:                  %s
		chan:                    %p
		pointer:                 %p

		%f     default width, default precision
		%9f    width 9, default precision
		%.2f   default width, precision 2
		%9.2f  width 9, precision 2
		%9.f   width 9, precision 0
	*/

	var message = "Hello World!"
	var d1, e1, f1 = 1, false, 3

	fmt.Println(message, d1, e1, f1)
}
