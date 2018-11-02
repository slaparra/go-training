package main

import (
	"fmt"
	"github.com/slaparra/go-training/src/02-package/how-to-use-a-package/aPackage"
)

func main() {
	fmt.Printf("Main file\n")
	aPackage.AMethod()
	aPackage.VisibleMethod()
}
