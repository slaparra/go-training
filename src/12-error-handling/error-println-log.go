package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	//initialize a log.txt file to insert the log information
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt") //try to find a non existing txt file to force an error
	if err != nil {
		//fmt.Println("err happened", err)
		//log.Println("err happened", err)
		//log.Fatalln(err)
		panic(err)
	}
}

/*
 Package log implements a simple logging package ... writes to standard error and prints the date and time of each
 logged message ... the Fatal functions call os.Exit(1) after writing the log message ...
 the Panic functions call panic after writing the log message.
*/

// log.Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.

// Fatalln is equivalent to log.Println() followed by a call to os.Exit(1).

// Panic shows an error message and the error line and info
