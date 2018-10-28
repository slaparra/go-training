package main

import (
	"fmt"
	"sort"
)

type training struct {
	hours int
	name  string
}

type myTraining []training

func main() {
	justDoIt := myTraining{
		{300, "AI"},
		{5, "SOLID"},
		{50, "DDD"},
		{200, "Event-Sourcing"},
		{100, "CQRS"},
	}

	isSortedMyTraining := sort.SliceIsSorted(justDoIt, func(i, j int) bool { return justDoIt[i].hours < justDoIt[j].hours })
	fmt.Println(justDoIt)           //[{300 AI} {5 SOLID} {50 DDD} {200 Event-Sourcing} {100 CQRS}]
	fmt.Println(isSortedMyTraining) //false

	sort.Sort(justDoIt)

	isSortedMyTraining = sort.SliceIsSorted(justDoIt, func(i, j int) bool { return justDoIt[i].hours < justDoIt[j].hours })
	fmt.Println(justDoIt)           //[{5 SOLID} {50 DDD} {100 CQRS} {200 Event-Sourcing} {300 AI}]
	fmt.Println(isSortedMyTraining) //true
}

func (mt myTraining) Len() int {
	return len(mt)
}

func (mt myTraining) Less(i, j int) bool {
	return mt[i].hours < mt[j].hours
}

func (mt myTraining) Swap(i, j int) {
	mt[i], mt[j] = mt[j], mt[i]
}
