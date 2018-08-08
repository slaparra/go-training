# Functions

- [Function examples](../todd-mcleod/06-functions)
- [Variadic](https://golang.org/ref/spec#Passing_arguments_to_..._parameters)


### Go functions and keywords

#### fmt.Sprint
```
// from package fmt: Sprint
// concatenate with space
fmt.Sprint('Santiago', 'Laparra')
```

#### append
```
// append: add an element to a slice
aSlice := {0, 1}
anotherSlice := append(aSlice, 2)

```

#### len: length of a slice
```
aVar := {1, 2}
fmt.Print(len(aVar)) // 2
```
[* length & capacity tour](https://tour.golang.org/moretypes/11)

#### range
Range iterates over elements in a variety of data structures
```
nums := []int{2, 3, 4}
sum := 0
for _, num := range nums {
    sum += num
}
fmt.Println("sum:", sum)
```
* https://github.com/golang/go/wiki/Range
* https://gobyexample.com/range

#### make
[Allocation with *make*](https://golang.org/doc/effective_go.html#allocation_make)
```
var s []byte

//create a slice len 5 cap 5
s = make([]byte, 5, 5)

// s == []byte{0, 0, 0, 0, 0}
```
[* from golang blog](https://blog.golang.org/go-slices-usage-and-internals)


#### Defer   
*Deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred.  
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