# Data structures

In this readme, we can find documentation and examples about:
- [Array](#array-definition-in-go)
- [Slice](#slice-definition)
- [Map](#map-definition-in-go)

## **Array definition in go**
[(Todd McLeod)](https://docs.google.com/document/d/1nt5bYAAS5sTVF6tpLaFLDHQzo5BNkcr4b507fg3ZPwM/edit#)
- *An array is a numbered sequence of elements of a single type.*
- *The number of elements is called the length and is never negative.*
- *The length is part of the array's type; it must evaluate to a non-negative constant representable by a value of type int.*
    - *The length of an array a can be discovered using the built-in function len.*
- *The elements can be addressed by integer indices 0 through len(a)-1.*
- *Array types are always one-dimensional but may be composed to form multi-dimensional types.*
- *Not dynamic, does not change in size*
- *The value property can be useful but also expensive; if you want C-like behavior and efficiency, you can pass a pointer to the array*
- *The first index is 0 position, so the last position is length-1*
- Links:
    - [Effective go](https://golang.org/doc/effective_go.html#arrays)
    - [Go spec](https://golang.org/ref/spec#Array_types)
    - [Tour golang](https://tour.golang.org/moretypes/6)

```
package main

import "fmt"

func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
}
```

## **Slice definition**  
[Golang spec](https://golang.org/ref/spec#Slice_types)  

[Todd McLeod:](https://docs.google.com/document/d/1nt5bYAAS5sTVF6tpLaFLDHQzo5BNkcr4b507fg3ZPwM/edit#)
- *A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array.*
- *The value of an uninitialized slice is nil.*
    - *it is a reference type*
- *Like arrays, slices are indexable and have a length.*
- *The length of a slice s can be discovered by the built-in function "len"*
    - *Unlike arrays, slices are dynamic*
        - *Their length may change during execution.*
- *The elements can be addressed by integer indices 0 through len(s)-1*
- *A slice, once initialized, is always associated with an underlying array that holds its elements.*
    - *it is a reference type*
- *The array underlying a slice may extend past the end of the slice.*
    - *Capacity is a measure of that extent:*
        - *it is the sum of the length of the slice and the length of the array beyond the slice;*
    - *The capacity of a slice a can be discovered using the built-in function cap(a).*
    - [Length and capacity](https://golang.org/ref/spec#Length_and_capacity)
- make
    - *A new, initialized slice value for a given element type T is made using the built-in function make, which takes a slice type and parameters specifying the length and optionally the capacity.*
    - *A slice created with make always allocates a new, hidden array to which the returned slice value refers.*
    ```
        make([]T, length, capacity)
        make([]int, 50, 100)            
        // same as this: new([100]int)[0:50]
    ```    
- *Like arrays, slices are always one-dimensional but may be composed to construct higher-dimensional objects. (multi-dimensional slices) *
- *Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.*)
- Links:
    - [Golang.org ref](https://golang.org/ref/spec#Slice_types)
    - [Slice reference type](../src/01-fundamentals/slice_reference_type.go)
    - [Value types and Reference types, The Way to Go](https://tinyurl.com/yah9vxcs)
    - [Slice usage and internals](https://blog.golang.org/go-slices-usage-and-internals)
    - [Slice expressions](https://golang.org/ref/spec#Slice_expressions)      
    - [Slice len & cap tour](https://tour.golang.org/moretypes/11)
    - [How to use capacity and length post](https://www.calhoun.io/how-to-use-slice-capacity-and-length-in-go/)


### Array & slice declaration
```
arrayLength5With2Elements := [5]string{"Hello ", "world"}
arrayLength2With2Elements = [...]string{"Penn", "Teller"}
length := len(arrayLength2With2Elements) //2

aSlice := []string{"Hello ", "world"}
```

String is a slice of bytes, string is made of runes, a rune is a unicode point.
```
fmt.Println("myString"[2])          //83
fmt.Println(string("myString"[2]))  //S
```

#### Slicing a slice
```
mySlice := []string{"a", "b", "c", "d", "e", "f", "g"}

//prints from position 2 to 4 (4 not included)
fmt.Print(mySlice[2:4])   //[c, d]  
```
[Delete an element from the slice, example](https://play.golang.org/p/MFmGqFGwW9i)

#### Append elements to a slice

```
package main

import (
    "fmt"
)

func main() {
    slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
    fmt.Printf("%v\n", slice)
    fmt.Printf("%T\n", slice)
    fmt.Println(len(slice))

    slice = append(slice, 52)       //append 52
    fmt.Printf("%v\n", slice)

    slice = append(slice, 53, 54)   //append 53, 54
    fmt.Printf("%v\n", slice)

    y:= []int{56,57,58}
    slice = append(slice[3:6], y...)     //return a slice from position 3 to 6 (6 not included)
    fmt.Printf("%v\n", slice)
    fmt.Println(slice[3:])     
}

// [42 43 44 45 46 47 48 49 50 51]
// []int
// 10
// [42 43 44 45 46 47 48 49 50 51 52]
// [42 43 44 45 46 47 48 49 50 51 52 53 54]
// [45 46 47 56 57 58]
// [56 57 58]
```

#### Make & new
```
// Both has the same result
make([]int, 50, 100)
new([100]int)[0:50]
```

[Todd McLeod slides: "slice, map, new, make, struct"](https://docs.google.com/presentation/d/1jot31JzJ7DiykCWpebfHz5_7s4JWZvklr-xmVWHHApU/edit#slide=id.gb91814ee3_0_17)    


#### Code examples
- [Array & slice](../src/07-data-structures/array_slice.go)  
- [Multidimensional slice](../src/07-data-structures/multi-dimensional-slice.go)  

## Map definition in go
*A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type*  
They are called dictionaries (key/value storage)  
- [Map type](https://golang.org/ref/spec#Map_types)
- [Maps, effective go](https://golang.org/doc/effective_go.html#maps)
- [Macro View of Map Internals In Go](https://www.ardanlabs.com/blog/2013/12/macro-view-of-map-internals-in-go.html)
- [Maps in action, Golang blog](https://blog.golang.org/go-maps-in-action)  
- ['Make' function to create a map](https://golang.org/ref/spec#Making_slices_maps_and_channels)
- [Delete a value of the map](https://golang.org/ref/spec#Deletion_of_map_elements)  
- [Golang book, Maps](https://www.golang-book.com/books/intro/6#section3)
- [**Map example](../src/07-data-structures/map.go)  

*Maps are Reference Types, they behave like pointers. When you pass a map variable to a function any changes to that mapped variable in the function change that original mapped variable outside the function*

*This looks very much like an array but there are a few differences.*
- *First the length of a map (found by doing len(x)) can change as we add new items to it. When first created it has a length of 0, after x[1] = 10 it has a length of 1.*
- *Second maps are not sequential. We have x[1], and with an array that would imply there must be an x[0], but maps don't have this requirement.*

To test for presence in the map without worrying about the actual value, you can use the blank identifier (_) in place of the usual variable for the value.

```
_, present := timeZone[tz]
```

##### Concurrency

[Maps are not safe for concurrent use](https://golang.org/doc/faq#atomic_maps): it's not defined what happens when you read and write to them simultaneously. 
If you need to read from and write to a map from concurrently executing goroutines, the accesses must be mediated by some kind of synchronization mechanism. 