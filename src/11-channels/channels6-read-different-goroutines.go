package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	ch := make(chan string, 1)
	go func() {
		ch <- "hi"
		wg.Done()
	}()

	go func() {
		value := <-ch
		value2 := <-ch
		fmt.Println(value)
		fmt.Println(value2)
		wg.Done()
	}()

	go func() {
		ch <- "hi2"
		wg.Done()
	}()
	wg.Wait()
}

//fatal error: all goroutines are asleep - deadlock!
