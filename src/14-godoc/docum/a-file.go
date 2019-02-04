// This package is documented with comments in a-file.go
package docum

import "fmt"

// Comment for the struct Person
type Person struct {
	First string
}

// Method Start is the main entry of the a-file.go application
func Start() {
	p := Person{"Santi"}
	p.Speak()
	fmt.Println("A comment")
	aMethod()
}

// This is a comment of the method speak
func (p *Person) Speak() {
	fmt.Println("Person speak")
}

// Commenting a method
func aMethod() {
	fmt.Println("aMethod")
}
