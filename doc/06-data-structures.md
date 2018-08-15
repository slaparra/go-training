# Data structures



- Array definition in go [(Todd McLeod)](https://docs.google.com/document/d/1nt5bYAAS5sTVF6tpLaFLDHQzo5BNkcr4b507fg3ZPwM/edit#)
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
        - [Own examples](../todd-mcleod/01-fundamentals/array_slice.go)       
- Slice
    - *The length can be changed* 
    - *Multi-dimensional*  
    - *Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.*)
    - Links:
        - [Golang.org ref](https://golang.org/ref/spec#Slice_types)
        - [Slice reference type](../todd-mcleod/01-fundamentals/slice_reference_type.go)
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