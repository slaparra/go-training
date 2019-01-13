package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	// wg.Add(2)        //uncomment this add to fix the race condition
	go func() {
		wg.Add(1) //remove this add to fix the race condition
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1) //remove this add to fix the race condition
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for number := range c {
		fmt.Println(number)
	}
}

//go run -race channels7-race-condition-waitgroup-add.go
