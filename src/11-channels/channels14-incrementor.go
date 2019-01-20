package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64

func main() {
	ch1 := incrementor2("1")
	ch2 := incrementor2("2")

	for n := range mergeChannels(ch1, ch2) {
		atomic.AddInt64(&count, int64(n))
	}

	fmt.Println("Final Counter:", count)
}

func incrementor2(s string) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("Process: "+s+" printing:", i)
			ch <- 1
		}
		close(ch)
	}()
	return ch
}

func mergeChannels(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
