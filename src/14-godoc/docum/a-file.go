// This package is documented with comments in a-file.go for go training: https://github.com/slaparra/go-training
package docum

import "fmt"

// Person Comment for the struct Person
type Person struct {
	First string
}

// Start Method is the main entry of the a-file.go application
func Start() string {
	p := Person{"Santi"}
	aMethod()
	return p.FirstName()
}

// First This is a comment of the method speak
func (p *Person) FirstName() string {
	return fmt.Sprint(p.First)
}

// aMethod Commenting a method
func aMethod() {
	fmt.Println("aMethod")
}
