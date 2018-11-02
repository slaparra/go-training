package main

import (
	"fmt"
	"time"
)

//from: https://tour.golang.org/concurrency/4

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		time.Sleep(time.Second * 1)
	}
	close(c)
}

func main() {
	c := make(chan int, 50)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
