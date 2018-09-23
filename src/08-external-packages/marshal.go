package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First               string
	Last                string `json:"Surname"` // with this notation we change the json key
	Age                 int
	notExported         int
	NotExportedNotation int `json:"-"` // with this notation we omit this value in the json
}

func main() {
	p1 := person{"James", "Bond", 20, 007, 8}
	fmt.Println(p1) //{James Bond 20 7 8}
	marshaledValue := marshal(&p1)
	fmt.Println(*marshaledValue) // ascii values of the json
	// [123 34 70 105 114 115 116 34 58 34 74 97 109 101 115 34 44 34 76 97 115 116 34 58 34 66 111 110 100 34 44 34 65 103 101 34 58 50 48 125]

	p2 := person{}
	unmarshal(marshaledValue, &p2)
	fmt.Println(p2) //{James Bond 20 0 0}
}

func marshal(p1 *person) *[]uint8 {
	bs, _ := json.Marshal(*p1)
	fmt.Printf("%T \n", bs) // []uint8
	fmt.Println(string(bs)) // {"First":"James","Last":"Bond","Age":20}
	// (JSON value)
	// NOTICE: notExported field is not in the json because is a not visible field

	return &bs
}

func unmarshal(marshaledValue *[]uint8, p *person) {
	json.Unmarshal(*marshaledValue, &p)
}
