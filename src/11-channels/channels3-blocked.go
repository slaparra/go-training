package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan string)
	go func() {
		ch <- "hi"
		ch <- "hi2"
		wg.Done()
	}()

	go func() {
		value := <-ch
		//value2 := <-ch
		fmt.Println(value)
		//fmt.Println(value2)
		wg.Done()
	}()
	wg.Wait()
}

//fatal error: all goroutines are asleep - deadlock!
