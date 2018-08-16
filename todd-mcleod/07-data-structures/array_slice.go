package main

import "fmt"

func main() {
	arrayFunc()
	sliceFunc()
	makeFunc()
	newFunc()
	sliceAnArray()
}

func arrayFunc() {
	var array4Pos [4]int //initialize an array of int, length 4 and default 0 value initialized [0 0 0 0]
	array4Pos[0] = 1
	i := array4Pos[0]
	fmt.Println(i) // i == 1

	//array4Pos[4] = 3 //error

	array2Elements := [2]string{"Penn", "Teller"}

	fmt.Printf("%T %v\n", array2Elements, array2Elements)

	array2Elements = [...]string{"Penn", "Teller"}

	fmt.Printf("%T %v\n", array2Elements, array2Elements)

	var x [58]string
	fmt.Println(x) //prints empty array initialized with "" => [                                                         ]

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	x[33] = "slaparra"
	fmt.Println(x)
	fmt.Println(len(x))
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

	//t is a new slice containing the same array of s but with a different length
	t := s[:4]
	printSlice(t)
	printSlice(s)

	//delete element position 2
	s2 := []int{2, 3, 5, 7, 11, 13}
	s2 = append(s2[:2], s2[3:]...)
	fmt.Println(s2)

	/*
		len=6 cap=6 [2 3 5 7 11 13]
		len=1 cap=6 [2]
		len=4 cap=6 [2 3 5 7]
		len=2 cap=4 [5 7]
		len=4 cap=4 [5 7 11 13]
		len=2 cap=4 [5 7]
	*/
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// slices are reference types and although 's' is passed by value, the slice s has the same reference as
	// the variable before the printSlice call. Has the same reference to the same array
	// so if we change s[0] inside this method, it will modify the slice in this scope and in the outer scope
	// https://golang.org/doc/effective_go.html#slices
}

func makeFunc() {
	var s []int
	s = make([]int, 5, 5)
	fmt.Printf("%T %v\n", s, s)

	slice := make([]int, 5, 100)
	// error out of bounds
	// slice[20] = 3
	slice[3] = 33
}

func newFunc() {
	aSlicePointer := new([]int)
	fmt.Println(aSlicePointer) // &[]

	anArrayPointer := new([3]int)
	fmt.Println(anArrayPointer) // &[0 0 0]

	aSliceUsingSlicing := new([3]int)[0:2]
	fmt.Println(aSliceUsingSlicing) // [0 0]
}

func sliceAnArray() {
	anArray := [30]string{}
	aSliceFromArraySliced := anArray[:]
	fmt.Println(aSliceFromArraySliced)
}
