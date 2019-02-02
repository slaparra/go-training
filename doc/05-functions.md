# Functions

- [Function examples](../src/06-functions)
- [Variadic](https://golang.org/ref/spec#Passing_arguments_to_..._parameters)

## Func expression

### Anonymous function
An anonymous function (function literal, lambda abstraction, or lambda expression) 
is a function definition that is not bound to an identifier (wiki)
```
func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
}
```
- [Example](../src/06-functions/function-variable.go)

![img](../resources/anonymous-gopher.png)  
<sub>(Image: [Trevor Forrey](https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3))</sub>

### Closure

A Closure is a technique for implementing lexically scoped name binding 
in a language with first-class functions. Operationally, a closure is a 
record storing a function together with an environment. (wiki)

```
func main() {
    newInts := intSeq()         
    fmt.Println(nextInt())      // 1
    fmt.Println(nextInt())      // 2
    fmt.Println(nextInt())      // 3

    newInts := intSeq()
    fmt.Println(newInts())      // 1
}

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```
\* [5 ways to use closure](https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/)  
\* [A Tour of go, closures](https://tour.golang.org/moretypes/25)


### Function as argument (callback)
```
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {	
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

```

- [Example](../src/06-functions/callback.go)

## Go functions and keywords

#### fmt.Sprint
```
// from package fmt: Sprint
// concatenate with space
fmt.Sprint('Santiago', 'Laparra')
```

#### append
```
// append: add an element to a slice
aSlice := []int{0, 1}
anotherSlice := append(aSlice, 2)

sliceToAppend := []int{3, 4}
theLastSlice := append(anotherSlice, sliceToAppend...)  //variadic to use append with a slice

fmt.Println(theLastSlice)   //[0 1 2 3 4] 
```

#### len: length of an array or slice
```
aVar := {1, 2}
fmt.Print(len(aVar)) // 2
```
[* length & capacity tour](https://tour.golang.org/moretypes/11)


#### make
[Allocation with *make*](https://golang.org/doc/effective_go.html#allocation_make)
```
var s []byte

//create a slice len 5 cap 5
s = make([]byte, 5, 5)

// s == []byte{0, 0, 0, 0, 0}
```
[* from golang blog](https://blog.golang.org/go-slices-usage-and-internals)


#### New
*new(T) allocates zeroed storage for a new item of type T and returns its address, a value of type *T. In Go terminology, it returns a pointer to a newly allocated zero value of type T.*  
[Allocation with *new*](https://golang.org/doc/effective_go.html#allocation_new)
```
aSlicePointer := new([]int) 
fmt.Println(aSlicePointer)              // &[]

anArrayPointer := new([3]int) 
fmt.Println(anArrayPointer)             // &[0 0 0]

aSliceUsingSlicing := new([3]int)[0:2] 
fmt.Println(aSliceUsingSlicing)         // [0 0]

```

#### Defer   

*Deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred.  

```
package main

import (
	"fmt"
)

func main() {
	defer myFunc1()
	myFunc2()
	fmt.Println("bla3")
}

func myFunc1() {
	fmt.Println("bla1")
}

func myFunc2() {
	fmt.Println("bla2")
}

// bla2
// bla3
// bla1
```

Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns.  
Deferred functions are executed in LIFO order.*

- https://golang.org/ref/spec#Defer_statements
- https://golang.org/doc/effective_go.html#defer

```

func main() {
	b()
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}



//    entering b
//    in b
//    leaving b

// defer executes "un" method before b() finishes 
// but its arguments (trace("b")) are evaluated immediately
```
*"This is useful for many reasons, the most common of which are to close an open connection or unlock a Mutex immediately before the function ends."* [(kylewbanks.com)](https://kylewbanks.com/blog/when-to-use-defer-in-go)  

*"A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions."* [(golang blog)](https://blog.golang.org/defer-panic-and-recover)

*Defer is a language mechanism that puts your function call into a stack. The deferred functions will be executed in reverse order when the host function finishes regardless of whether a panic is called or not. The defer mechanism is very useful for cleaning up resources.* [(@hussachai)]

[(@hussachai)]: https://medium.com/@hussachai/error-handling-in-go-a-quick-opinionated-guide-9199dd7c7f76

#### Init function
*init() is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized.*  
```
func init() {
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath may be overridden by --gopath flag on command line.
    flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}
```
https://golang.org/doc/effective_go.html#init  
https://golang.org/ref/spec#Package_initialization

### Printing
If we implement the method String for a type it will be used when fmt.print methods are executed  

*If you want to control the default format for a custom type, all that's required is to define a method with the signature String() string on the type.* 
[(effective go)](https://golang.org/doc/effective_go.html#printing)

```

type people []string

studyGroup := people{"Zeno", "John", "Al", "Jenny"}

func (p people) String() string {
	return fmt.Sprintf("- Custom people type print String [%s, %s, %s, %s]", p[0], p[1], p[2], p[3])
}
```
