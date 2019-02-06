package main

import (
	"13-testing/04-benchmark/somePackage"
	"fmt"
)

func main() {
	s := "Some text"
	fmt.Println(somePackage.PrintAString(s))
}
