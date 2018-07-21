package aPackage

import "fmt"

func AMethod() {
	fmt.Printf("A visible method in aPackage\n")
	fmt.Printf("Print a var from another file: '%s'\n", aString)
	anotherMethod()
}
