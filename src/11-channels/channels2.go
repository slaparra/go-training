package main

//example from https://pragmacoders.com/multithreading-go-tutorial/

import (
	"fmt"
)

func main() {
	a := 1
	b := 2

	operationDone := make(chan bool) //Make a channel that consumes and outputs bool values.
	go func() {
		b = a * b

		operationDone <- true
	}()

	<-operationDone //the channel waits until a message is received (sent at the end of the anonymous function)

	a = b * b

	fmt.Println("Hit Enter when you want to see the answer")
	fmt.Scanln()

	fmt.Printf("a = %d, b = %d\n", a, b)
}

/*
   with channels, the result is:
   a = 4, b = 2

   without channels, the result is:
   a = 4, b = 8
   the anonymous function is called at the end
*/
