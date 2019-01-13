package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	fmt.Println(<-c)
}

// Why does this only print zero?
// And what can you do to get it to print all 0 - 9 numbers?

/*
    solution: use a "for - range" to get all the elements in the channel buffer
			  because println only get the first one element.
              Don't forget to close the channel

    func main() {
        c := make(chan int)

        go func() {
            for i := 0; i < 10; i++ {
                c <- i
            }
->          close(c)
        }()

->      for n := range c {
->          fmt.Println(n)
->      }
    }
*/
