package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
}

/*
at the terminal:
go run os.go <your name>
*/
