package aPackage

import "fmt"

//this method is only visible in the go files of aPackage package
func anotherMethod() {
	fmt.Printf("another-file.go - anotherMethod - Another method not visible for main in aPackage. Called from some-file.go\n")
}

//the method is visible outside the package if the name is in capital letter
func VisibleMethod() {
	fmt.Printf("another-file.go - VisibleMethod - Another method visible for main in aPackage. Called from main.go\n")
}
