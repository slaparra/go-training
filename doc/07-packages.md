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
