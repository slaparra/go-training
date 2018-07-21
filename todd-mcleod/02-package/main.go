package main

import (
	"github.com/slaparra/go-training/todd-mcleod/02-package/aPackage"
	"fmt"
)

func main() {
	fmt.Printf("Main file\n")
	aPackage.AMethod()
	aPackage.VisibleMethod()
}
