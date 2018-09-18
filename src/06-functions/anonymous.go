package main

import "fmt"

var name string = ""

func main() {
	anonymous()
	closureFunc()
}

func anonymous() {
	anonyFunction := func(name string) string {
		return "Hi1 " + name + "!!"
	}

	fmt.Print("Enter a name: ")
	fmt.Scan(&name)

	fmt.Println(anonyFunction(name))
}

func closureFunc() {
	aFunc := func(a string) string {
		return "Hi2 " + a + "!!"
	}

	fmt.Println(aFuncWithAFuncAsArgument(aFunc))
}

func aFuncWithAFuncAsArgument(f func(string) string) string {
	fmt.Print("Enter a name: ")
	var aVar string
	fmt.Scan(&aVar)

	return f(aVar)
}
