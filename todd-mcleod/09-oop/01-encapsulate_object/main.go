package main

import (
	"fmt"
	"github.com/slaparra/go-training/todd-mcleod/09-oop/01-encapsulate_object/author"
)

// go run main.go author
func main() {
	aut := author.Create("Santiago", "Laparra", "My bio example")

	fmt.Println(aut.FullName())

	fmt.Println(aut.Bio())
}
