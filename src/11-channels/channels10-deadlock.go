package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

/*
    solution to prevent deadlock is evolving with a goroutine

    func main() {
        c:= make(chan int)
->      go func() {
->          c <- 1
->      }()
        fmt.Println(<-c)
    }

*/
