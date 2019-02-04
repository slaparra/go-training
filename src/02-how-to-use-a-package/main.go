// This subject shows how to build a package with different files inside
package main

import (
	"fmt"
	"github.com/slaparra/go-training/src/02-how-to-use-a-package/aPackage"
)

func main() {
	fmt.Printf("Main file\n")
	aPackage.AMethod()
	aPackage.VisibleMethod()
}
