# Go packages

- [Packages (book An Introduction to Programming in Go )](https://www.golang-book.com/books/intro/11)
- [Package file examples](../src/02-package)
- Methods and vars have to be capitalized to be visible outside the package


```
# when a file needs to access to a struct from another in the same package 
# we need to run "go install" or "run" the different files:
go run main.go other_file.go

# another package
go run main.go aPackage
```

## External packages
- [Bufio newscanner](https://golang.org/pkg/bufio)
- [Net/http](https://golang.org/pkg/net/http/) 
- [Enconding/json](https://godoc.org/encoding/json)  
- [Runtime](https://golang.org/pkg/runtime/)
- [Sort](https://golang.org/pkg/sort/)
- [Sync](https://golang.org/pkg/sync/)
- [Time](https://golang.org/pkg/time/)
- [Math/rand](https://golang.org/pkg/math/rand/)
- [Bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt)

[Go package example files](../src/08-external-packages/)

### Examples

#### NumCPU and GOMAXPROCS
```
import (
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
```
- GOMAXPROCS sets the maximum number of CPUs that can be executing  
- NumCPU returns the number of logical CPUs usable by the current process.
As Go 1.5 Release Notes says

By default, Go programs run with GOMAXPROCS set to the number of cores available; in prior releases it defaulted to 1.

So starting from Go 1.5, the default value should be the number of cores.

#### Pkg sync, WaitGroup
*A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.*

```
import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup

func main() {
    wg.Add(2)   // if we remove WaitGroup feature, the application will end without finish foo & bar execution
    go foo()    // (because go routine)
    go bar()
    wg.Wait()   // It will wait until wg.Done() is executed twice
}

func foo() {
    for i := 0; i < 45; i++ {
        fmt.Println("Foo:", i)
    }
    wg.Done()
}

...
```
- [wait-group.go, Todd McLeod example](../src/08-external-packages/wait-group.go)
- [golang.org/pkg WaitGroup example](https://golang.org/pkg/sync/#example_WaitGroup)

#### Pkg Time, sleep

*Sleep pauses the current goroutine for at least the duration d. A negative or zero duration causes Sleep to return immediately.*

```
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Second)
	}
}	
```

- [Time, golang by example](https://gobyexample.com/time)

#### Time formatting

*Go supports time formatting and parsing via pattern-based layouts.*

- [Time formatting example**](../src/01-fundamentals/time-formating.go)
- https://gobyexample.com/time-formatting-parsing

#### Time add & substract

- [File Example](../src/01-fundamentals/time-add-substract.go)

- https://stackoverflow.com/questions/26285735/subtracting-time-duration-from-time-in-go)
- https://www.socketloop.com/tutorials/golang-minus-time-with-time-add-or-time-adddate-functions-to-calculate-past-date

#### Bcrypt

*Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.*

```
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(`aPassword`), 8)

...

if err = bcrypt.CompareHashAndPassword([]byte(`aPassword`), []byte(creds.Password)); err != nil {
    // If the two passwords don't match, return a 401 status
    fmt.Println("Invalid password")
}
```

- https://godoc.org/golang.org/x/crypto/bcrypt
- https://www.sohamkamani.com/blog/2018/02/25/golang-password-authentication-and-storage/

## Links
- [Parsing json in golang](https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/)
- [Json, GoByExample](https://gobyexample.com/json)