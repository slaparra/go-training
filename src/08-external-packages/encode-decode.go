package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type aPerson struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	encode()
	decode()
}

func encode() {
	p1 := aPerson{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1) //encode and write to the screen (os.Stdout follows Writer interface)
}

func decode() {
	var p1 aPerson
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
