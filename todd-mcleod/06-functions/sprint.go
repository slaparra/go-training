package main

import "fmt"

func main() {
	aString := callSprint("Santiago ", "Laparra") //concat
	fmt.Println(aString)
}

func callSprint(s string, s2 string) (response string) {
	response = fmt.Sprint(s, s2)
	return
}
