package main

import "fmt"

func main() {
	arrayFunc()
	sliceFunc()
	makeFunc()
}

func arrayFunc() {
	var array4Pos [4]int
	array4Pos[0] = 1
	i := array4Pos[0]
	fmt.Println(i) // i == 1

	//array4Pos[4] = 3 //error

	array2Elements := [2]string{"Penn", "Teller"}

	fmt.Printf("%T %v\n", array2Elements, array2Elements)

	array2Elements = [...]string{"Penn", "Teller"}

	fmt.Printf("%T %v\n", array2Elements, array2Elements)
}

//from https://tour.golang.org/moretypes/11
func sliceFunc() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the sliceFunc to give it zero length.
	s = s[:1]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	/*
		len=6 cap=6 [2 3 5 7 11 13]
		len=1 cap=6 [2]
		len=4 cap=6 [2 3 5 7]
		len=2 cap=4 [5 7]
	*/
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func makeFunc() {
	var s []int
	s = make([]int, 5, 5)
	fmt.Printf("%T %v\n", s, s)
}