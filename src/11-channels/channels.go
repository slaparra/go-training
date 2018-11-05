package main

//from: https://www.golang-book.com/books/intro/10

import (
	"fmt"
	"strconv"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- strconv.Itoa(i)
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong " + strconv.Itoa(i)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
