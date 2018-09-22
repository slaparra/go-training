package main

import (
	"fmt"
)

// from: Todd McLeod https://github.com/GoesToEleven/GolangTraining
// https://github.com/GoesToEleven/GolangTraining/blob/master/20_struct/05_promotion/02_overriding-methods/main.go

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz doubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {
	p1 := person{
		Name: "Ian Flemming",
		Age:  44,
	}

	p2 := doubleZero{
		person: person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}
	fmt.Println("File got from: Todd McLeod https://github.com/GoesToEleven/GolangTraining\n")

	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}

/*
Result:
	I'm just a regular person.
	Miss Moneypenny, so good to see you.
	I'm just a regular person.
*/
