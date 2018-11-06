package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	myChannel := make(chan string, 4)
	mySlice := []string{"hi", "how", "are", "you"}
	go func() {
		for _, item := range mySlice {
			myChannel <- item
		}
		close(myChannel)
		wg.Done()
	}()

	go func() {
		for myString := range myChannel {
			println(myString)
		}
		wg.Done()
	}()

	wg.Wait()
}

/*
   Because we closed the channel above, the iteration terminates after receiving the 2 elements.
   If we don't put close(mychanel) an error would be thrown: fatal error: all goroutines are asleep - deadlock!
*/
