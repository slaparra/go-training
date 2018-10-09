package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println(runtime.NumCPU())

	wg2.Add(2)
	go foo2()
	go bar2()
	wg2.Wait()
}

func foo2() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg2.Done()
}

func bar2() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg2.Done()
}
