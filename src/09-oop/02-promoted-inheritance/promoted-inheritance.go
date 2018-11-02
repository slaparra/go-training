package main

import (
	"fmt"
)

type potentialAdult interface {
	isOldEnough() bool
}

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person        //with this embedded we inherit all the values of person in doubleZero
	LicenseToKill bool
}

func (p *person) isOldEnough() bool {
	return p.Age > 18
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   16,
		},
		LicenseToKill: false,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	printIsAdult(&p1)
	printIsAdult(&p2)
}

func (d doubleZero) String() string {
	return fmt.Sprintf("First: %s, last: %s, age: %d, licenseToKill: %t	", d.First, d.Last, d.Age, d.LicenseToKill)
}

func printIsAdult(p potentialAdult) {
	fmt.Println(p.isOldEnough())
}
