package main

func main() {
	mySlice := []string{"hi", "how", "are", "you"}
	myChannel := make(chan string, len(mySlice)) //buffered with 4

	go func() {
		for _, item := range mySlice {
			myChannel <- item
		}
		close(myChannel)
	}()

	for myString := range myChannel {
		println(myString)
	}
}

/*
   Because we closed the channel above, the iteration terminates after receiving the 2 elements.
   If we don't put close(mychanel) an error would be thrown: fatal error: all goroutines are asleep - deadlock!
*/
