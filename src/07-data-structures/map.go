package main

import "fmt"

func main() {
	// shorthand creating a map
	aMapVar := map[string]int{
		"index1": 3,
		"index2": 5,
	}
	fmt.Printf("%T: %v\n", aMapVar, aMapVar)

	// adding an entry
	aMapVar["index3"] = 8
	aMapVar["index4"] = 14
	fmt.Printf("%T: %v\n", aMapVar, aMapVar)

	// creating a map with make,
	anotherMapVar := make(map[string]int, 100)
	fmt.Printf("%T: %v\n", anotherMapVar, anotherMapVar)

	//delete a key value of a map
	delete(aMapVar, "index2") // remove element m[k] from map m
	fmt.Printf("%T: %v\n", aMapVar, aMapVar)

	//the result of accessing to a position of the map is the value and a boolean (comma ok idiom)
	//this boolean is the result of if exists or not a value for the given key
	value, ok := aMapVar["index1"]
	fmt.Printf("value: %v, result: %v\n", value, ok) //value: 3, result: true

	fmt.Println("Map Range loop: ")
	mapForRange := map[string]string{
		"first-key":  "first-value",
		"second-key": "second-value",
		"third-key":  "third-value",
		"fourth-key": "fourth-value",
	}

	for key, value := range mapForRange {
		fmt.Printf("- key %s, value %s\n", key, value)
	}

	mapStringSlices := map[string][]string{
		"aKey":       []string{"1value", "2values"},
		"anotherKey": []string{"1value", "2values"},
	}

	fmt.Println(mapStringSlices)

	delete(mapStringSlices, "aKey")

	fmt.Println(mapStringSlices)
}
