package main

import "fmt"

func main() {

	//common loop
	for i := 32; i < 123; i++ {
		fmt.Printf("%d \t %b \t %x \t %q\n", i, i, i, i) //%q ascii
	}

	//like a while (while doesn't exists in Go)
	var c string
	for c != "continue" {
		fmt.Print("Type 'continue' to continue:")
		fmt.Scan(&c)
	}

	//like for (;;). Needs to do a break to end the loop
	var aVar int = 0
	for {
		if aVar == 50 {
			fmt.Println("It finish when aVar is equal to 50.")
			break
		}
		aVar++
	}

	//'continue' force a new iteration without execute the complete loop.
	aVar = 0
	for {
		aVar++
		if aVar < 50 {
			continue
		}

		fmt.Printf("This message should be printed once (aVar value %d)\n", aVar)
		break
	}

	//a loop with a range
	var aSlice = []int{1, 3, 5, 6}
	for key, value := range aSlice {
		fmt.Printf("key: %d, value: %d\n", key, value)
	}
}
