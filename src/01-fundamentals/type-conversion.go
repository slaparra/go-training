package main

import "fmt"

func main() {

	type person struct {
		name    string
		surname string
	}

	type people []person
	type crowd []person

	p1 := people{person{"Mad", "Max"}, person{"Johny", "Walker"}}
	fmt.Printf("%T: %v\n", p1, p1)

	//type conversion is possible because both are the same type slice of person
	m2 := crowd(p1)
	fmt.Printf("%T: %v\n", m2, m2)
}
