package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	exercice1()
	exercice2()
	exercice3()
	exercice4()
}

func exercice1() {
	fmt.Println("---- Exercice 1 ----")
	//The methods Less, Len & Swap are created for type people
	//with this way people implements the interface Interface. Now the slice can be sortable with sort.Sort
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}

func exercice2() {
	fmt.Println("---- Exercice 2 ----")
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.StringSlice(s).Sort()
	fmt.Println(s)

	p := people(s)
	sort.Sort(p)
	fmt.Println(p)
}

func exercice3() {
	fmt.Println("---- Exercice 3 ----")
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.IntSlice(n).Sort()
	fmt.Println(n)
}

func exercice4() {
	fmt.Println("---- Exercice 4 ----")
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

//https://stackoverflow.com/questions/38571354/best-way-to-swap-variable-values-in-go
func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Len() int {
	return len(p)
}

//if we implement the method String for a type
//it will be used when fmt.print methods are executed
func (p people) String() string {
	return fmt.Sprintf("- Custom people type print String [%s, %s, %s, %s]", p[0], p[1], p[2], p[3])
}
