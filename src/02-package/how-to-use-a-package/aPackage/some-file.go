package aPackage

import "fmt"

func AMethod() {
	fmt.Printf("some-file.go - AMethod - A visible method in aPackage. Called from main.go\n")
	fmt.Printf("some-file.go - AMethod - Print a var from another file: '%s'\n", aString)
	anotherMethod()
}
