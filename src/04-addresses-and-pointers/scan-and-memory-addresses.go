package main

import "fmt"

func main() {
	var aVar string = "Hello var"

	fmt.Printf("print aVar value: '%s'\n", aVar)
	fmt.Printf("print aVar address %#x\n", &aVar)

	fmt.Printf("Enter a new word for aVar: ")
	fmt.Scan(&aVar)

	fmt.Printf("You've typed: %s\n", aVar)
	fmt.Printf("and the value '%s' is saved in the same aVar variable address %#x\n", aVar, &aVar)
}

/*
	print aVar value: 'Hello var'
	print aVar address 0xc42000e1e0
	Enter a new word for aVar: helloWar
	You've typed: helloWar
	and the value 'helloWar' is saved in the same aVar variable address 0xc42000e1e0
*/
