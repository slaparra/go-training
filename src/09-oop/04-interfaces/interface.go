package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first     string
	last      string
	birthYear int
}

type dog struct {
	name      string
	birthYear int
}

func main() {
	me := person{"Santiago", "Laparra", 2010}
	toby := dog{"Toby", 2015}

	fmt.Println(me)
	fmt.Printf("Type of me: %T\n", me)
	me.speak()

	amIAHuman(me)

	// Toby is not a human so we can't call amIAHuman(toby). It fails because there's no method speak related with dog
	//
	// amIAHuman(toby)
	//
	// Error:
	// cannot use toby (type dog) as type human in argument to amIAHuman:
	//	dog does not implement human (missing speak method)

	toby.bark()

}

func (object person) speak() {
	fmt.Printf("My name is %s %s (%d)\n", object.first, object.last, object.birthYear)
}

func amIAHuman(aHuman human) {
	fmt.Printf("In this example, %v is a person and a human too because struct person implements speak like human interface\n", aHuman)
}

func (object dog) bark() {
	fmt.Printf("Only dogs like <%s> can bark\n", object.name)
}
