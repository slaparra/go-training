package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)
	n := 10

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}

			done <- true
		}()
	}

	go func() {
		for i := 0; i < 10; i++ {
			<-done
		}
		close(c)
	}()

	for element := range c {
		fmt.Println(element)
	}
}
