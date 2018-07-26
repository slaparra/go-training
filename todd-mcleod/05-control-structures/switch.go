package main

import "fmt"

type square struct {
	length float32
	width  float32
}

func main() {

	defaultSwitch()

	switchWithFallThrough()

	switchWithTwoConditionsInOneCase()

	switchWithoutCondition()


	var aSquare = square{3.3, 4.5}
	switchTypeCondition(aSquare)
	switchTypeCondition(aSquare.width)
	switchTypeCondition(3)
	switchTypeCondition("a string")
}

func defaultSwitch() {
	var word string

	fmt.Print("Write 'one', 'two' or other word:")
	fmt.Scan(&word)

	switch word {
	case "one":
		fmt.Println("You wrote one")
	case "two":
		fmt.Println("You wrote two")
	default:
		fmt.Println("You wrote another word")
	}
}

func switchWithFallThrough() {
	var word string

	fmt.Print("Write 'two' to see fallthrough in action:")
	fmt.Scan(&word)

	switch word {
	case "one":
		fmt.Println("You wrote one")
	case "two":
		fmt.Println("You wrote two with optionFallThrough print three")
		fallthrough
	case "three":
		fmt.Println("You wrote three because the fallthrough")
	case "four":
		fmt.Println("You wrote four")
	default:
		fmt.Println("You wrote another word")
	}
}

func switchWithTwoConditionsInOneCase() {
	var word string

	fmt.Print("Write 'four' or 'five' to see 2 conditionals in one case:")
	fmt.Scan(&word)

	switch word {
	case "one":
		fmt.Println("You wrote one")
	case "two":
		fmt.Println("You wrote two")
	case "three":
		fmt.Println("You wrote three")
	case "four", "five":
		fmt.Println("You enter here if you type the option four or five")
	default:
		fmt.Println("You wrote another word")
	}
}

func switchWithoutCondition() {
	var aNumber int

	fmt.Print("Write a number: ")
	fmt.Scan(&aNumber)

	switch {
	case aNumber%2 == 1:
		fmt.Printf("%d is odd\n", aNumber)
	case aNumber%2 == 0:
		fmt.Printf("%d is even\n", aNumber)
	default:
		fmt.Println("may be an err?")
	}
}

func switchTypeCondition(aType interface{}) {

	fmt.Printf("%v ", aType)

	switch aType.(type) {
	case square:
		fmt.Println("Is a square type")
	case int:
		fmt.Println("Is int type")
	case float32:
		fmt.Println("Is float type")
	case string:
		fmt.Println("Is string type")
	default:
		fmt.Println("You used another type")
	}
}
