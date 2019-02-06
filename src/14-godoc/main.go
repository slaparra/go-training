// The package main doc. For go training: https://github.com/slaparra/go-training
package main

import (
	"fmt"
	"github.com/slaparra/go-training/src/14-godoc/docum"
)

// The main function
func main() {
	fmt.Println(docum.Start())
}

/*
   Type in the console:
       - go doc
       - go doc docum
       - go doc docum Person
       - go doc docum Person.First
       - go doc docum Start (or go doc docum.Start)

       - go doc fmt
       - go doc fmt.Errorf

   To see the documentation
*/
