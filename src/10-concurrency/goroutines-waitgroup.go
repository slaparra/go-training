package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(2)

	fmt.Println("Max CPUs:", runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	fmt.Println("CPUs:", runtime.cp())
	fmt.Println("Starting...")

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("do anotherthing")
		wait.Done()
	}()

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("do something")
		wait.Done()
	}()

	fmt.Println("Goroutines in process:", runtime.NumGoroutine())

	wait.Wait()

}

//without Waitgroup the app with goroutine is unpredictable and (without sleep) most of the time it will print only "Starting..."
