package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}
type people []person

func main() {
	aPerson := person{"John", "James", 23}
	anotherPerson := person{"Mariah", "Robin", 25}

	crowd := people{aPerson, anotherPerson}

	fmt.Println(crowd)
	fmt.Printf("%v %T", crowd, crowd)
}

func (p person) String() string {
	return fmt.Sprintf("{Name: %s, Surname: %s, Age: %d}", p.name, p.surname, p.age)
}
