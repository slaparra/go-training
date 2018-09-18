package main

import "fmt"

func main() {
	fmt.Println(43) //print line
	//print with format
	fmt.Printf("%d - %b - %s\n", 43, 43, "github.com/slaparra") //decimal, binary, string
	fmt.Printf("%d \t %b \t %f\n", 27, 27, 27.2)                //tabs. %f float
	fmt.Printf("%x \t %#x \t %#X\n", 27, 27, 27)                //hexadecimal
}
