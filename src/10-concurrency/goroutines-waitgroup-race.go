package main

import (
	"fmt"
	"runtime"
	"sync"
)

var incrementor int

var wait sync.WaitGroup
var locker sync.Mutex

func main() {

	gr := 100
	wait.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			fmt.Println("Goroutines:", runtime.NumGoroutine())
			locker.Lock()

			a := incrementor
			a++
			incrementor = a
			fmt.Println("Incrementor", incrementor) // putting this println will cause a race condition because more
			// than one goroutine can read the same variable
			locker.Unlock()
			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println("Incrementor", incrementor)
}
