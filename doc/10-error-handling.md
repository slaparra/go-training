# Error handling

*Go doesn’t have exceptions, so it doesn’t have try…catch or anything similar. How can we handle errors in Go then?*

*There are two common methods for handling errors in Go — **Multiple return values and panic**.* ([@hussachai])

[@hussachai]: https://medium.com/@hussachai/error-handling-in-go-a-quick-opinionated-guide-9199dd7c7f76

*In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C.*

## Create an error

#### errors.New

`func New(text string) error`

New returns an error that formats as the given text.

```
func main() {  
    _, err := Sqrt(-1)
    if err != nil {
        fmt.Println(err)
    }
    //...
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // implementation
}
```

#### fmt.Errorf

*The fmt package's Errorf function lets us use the package's formatting features to create descriptive error messages.*

```
package main

import (
    "fmt"
)

func main() {
    const name, id = "bimmler", 17
    err := fmt.Errorf("user %q (id %d) not found", name, id)
    if err != nil {
        fmt.Print(err)
    }
}
```
#### Custom errors

*Sometimes the caller needs extra context in order to make a more informed error handling decision. For me, that is when custom error types make sense.* 

```
// Naming convention for custom error types:
// --> It is idiomatic in Go to postfix the name of a custom error type with the word Error.
type CustomError struct {
    param    int
    errorMsg string
    method   string
}

// Error() to follow error interface: https://golang.org/pkg/errors/
func (error *CustomError) Error() string {
    return fmt.Sprintf("Error: %d %s - Method: %s", error.param, error.errorMsg, error.method)
}
```
- [Custom error example](../src/12-error-handling/custom-error.go)
- [Custom errors, golangbot](https://golangbot.com/custom-errors/)
- [How to build custom errors, Ardanlabs](https://www.ardanlabs.com/blog/2014/11/error-handling-in-go-part-ii.html)


## Handling errors

```
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f
```


There are 4 forms of handling errors:
```
_, err := os.Open("no-file.txt")    //try to find a non existing txt file to force an error
if err != nil {
    // fmt.Println("err happened", err)
    // log.Println("err happened", err)
    // log.Fatalln(err)
    // panic(err)
}
```
- *`log.Println` calls Output to print to the standard logger. Arguments are handled in the manner of `fmt.Println`.*
- *`log.Fatalln` is equivalent to `log.Println()` followed by a call to `os.Exit(1)`.*  [Exit into](https://godoc.org/os#Exit)
- [Error handling, golangbot](https://golangbot.com/error-handling/)

---

### Panic

[Panic] shows an error message,the error line and `exit status`. The panic built-in function stops normal execution of the current goroutine.

[Panic]: https://godoc.org/builtin#panic

```
panic: open no-file.txt: no such file or directory

goroutine 1 [running]:
main.main()
    /Users/santiago/Documents/Workspace/go/src/github.com/slaparra/go-training/src/12-error-handling/error-println-log.go:24 +0x63
exit status 2
```

This termination sequence is called panicking and can be controlled by the built-in function recover.

Panic is a built-in function that stops the normal execution flow. **The deferred functions are still run as usual.**

When you call panic and you don’t handle it, the execution flow stops, all deferred functions are executed in reverse order, and stack traces are printed at the end.

*One important factor is that you should avoid panic and recover and use errors where ever possible. Only in cases where the program just cannot continue execution should a panic and recover mechanism be used.*

- [Example: Defer is executed regardless a panic](https://play.golang.org/p/sfkGfBo04d3)  

### Recover
Recover is a built-in function that regains control of a panicking goroutine.

Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

```
package main

import (
    "errors"
    "fmt"
)

func main() {
    callFunc()

    fmt.Println("Returned normally from callFunc.")
}

func callFunc() {
    //if a method can result in a panic:
    //  put a defer inside a function
    //  and call the recover method to don't break the execution
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from defer function: ", r)
        }
    }()

    //force a panic (imagine a call to an external library)
    panic(errors.New("the error was panicked"))
}

```
Recover works only when it is called from the same goroutine. It's not possible to recover from a panic that has happened in a different goroutine. Let's understand this using an example.

- [Recovering example](../src/12-error-handling/recover.go). From: [Golang blog](https://blog.golang.org/defer-panic-and-recover)
- [Golangbot, Panic and Recover](https://golangbot.com/panic-and-recover/)
---

#### Examples
- [Create an error](../src/12-error-handling/create-error.go)
- [Custom errors](../src/12-error-handling/custom-error.go)
- [Bufio error variables](https://golang.org/src/bufio/bufio.go)
- [Show errors](../src/12-error-handling/error-println-log.go)


#### Links
- https://golang.org/ref/spec#Errors
- [**Error Handling in Go that Every Beginner should Know**](https://medium.com/@hussachai/error-handling-in-go-a-quick-opinionated-guide-9199dd7c7f76)
- [GoBlog: Error handling and go](https://blog.golang.org/error-handling-and-go)
- [GoByExample: Errors](https://gobyexample.com/errors)
- [Custom errors](https://golangbot.com/custom-errors/)
- [GoByExample: Panic](https://gobyexample.com/panic)
- [Panic & recover](https://golangbot.com/panic-and-recover/)
- [Defer, panic & recover](https://www.golang-book.com/books/intro/7#section6)

#### Videos
- [Don't Just Check Errors Handle Them Gracefully, Dave Cheney](https://www.youtube.com/watch?v=lsBF58Q-DnY)
