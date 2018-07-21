package aPackage

import "fmt"

//this method is only visible in the go files of aPackage package
func anotherMethod() {
	fmt.Printf("Another method not visible for main in aPackage\n")
}

//the method is visible outside the package if the name is in capital letter
func VisibleMethod() {
	fmt.Printf("Another method visible for main in aPackage\n")
}
