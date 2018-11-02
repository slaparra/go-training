package main

// example from -> https://github.com/GoesToEleven/GolangTraining

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2) // if we remove WaitGroup feature, the application will end without finish foo & bar execution
	go foo()  // (because go routine)
	go bar()
	wg.Wait() // It will wait until wg.Done() is executed twice
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
