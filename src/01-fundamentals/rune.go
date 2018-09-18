package main

import "fmt"

func main() {
	var aRune, anotherRune rune = 'i', 'z'

	//https://en.wikipedia.org/wiki/ASCII#Printable_characters
	fmt.Printf("rune value %d, type: %T, casted to string: %s\n", aRune, rune(aRune), string(aRune))
	fmt.Printf("anotherRune value %d, type: %T, casted to string: %s\n", rune(anotherRune), anotherRune, string(anotherRune))
	//cast to rune is not needed

	//var aString string = "a"
	//fmt.Printf("%v", rune(aString))
	//a string can't be converted to rune

	for i := 103; i <= 109; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
	for i := 2000; i <= 2005; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
