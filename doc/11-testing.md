# Testing

Go includes a special package ([testing]) that makes writing tests easier.

[testing]:https://golang.org/pkg/testing/

Characteristics of a Golang test function:

- The first and only parameter needs to be `t *testing.T`
- Be in a function `func TestXXX(t *testing.T)`.
- It begins with the word `Test` followed by a word or phrase starting with a capital letter.
(usually the method under test i.e. `TestValidateClient`)
- Calls `t.Error` or `t.Fail` to indicate a failure (I called `t.Errorf` to provide more details)
- `t.Log` can be used to provide non-failing debug information
- Must be saved in a file named `something_test.go` such as: `addition_test.go`


Run tests:
`go test` or `go test -v` to run tests with verbose mode

## Testing package

*Package [testing] provides support for automated testing of Go packages. It is intended to 
be used in concert with the `go test` command, which automates execution of any function of the form.*

[testing]: https://golang.org/pkg/testing/

```
package math

import "testing"

func TestAverage(t *testing.T) {
  var v float64
  v = Average([]float64{1,2})
  if v != 1.5 {
    t.Error("Expected 1.5, got ", v)
  }
}

---

$ go test
PASS
ok      golang-book/chapter11/math      0.032s
```
`t.Error` is equivalent to `t.Log followed by Fail.
- https://golang.org/pkg/testing/
- https://godoc.org/testing

## Table driven tests (Test tables)
*At the heart of all table driven tests is the table itself, which provides the inputs and expected results of the function under test. In most cases the table is a slice of anonymous structs, which allows the table to be written in a compact form.* ([dave-cheney])

[dave-cheney]: https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go

#### Example:
Method to test
```
package fib

// Fib returns the nth number in the Fibonacci series.
func Fib(n int) int {
        if n < 2 {
                return n
        }
        return Fib(n-1) + Fib(n-2)
}
```

Test file:
```
type fibTest struct {
        n        int
        expected int
}

var fibTests = []fibTest {
        {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13},
}

func TestFib(t *testing.T) {
        for _, tt := range fibTests {
                actual := Fib(tt.n)
                if actual != tt.expected {
                        t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
                }
        }
}
```
*The use of t.Errorf instead of t.Fatalf is a personal preference. As Fib is a pure function it is safe to continue the loop after a failure. I find this generally reduces test whack-a-mole by returning all the failures at once.*

- [Factorial with table driven test example](../src/13-testing/02-table-driven-test/factorial_test.go)
- [Table driven test wiki golang](https://github.com/golang/go/wiki/TableDrivenTests)

### Generating an HTML coverage report

If you use the following two commands you can visualise which parts of your program have been covered by the tests and which statements are lacking:

```
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html 
```

Then open coverage.html in a web-browser ([blog.alexellis.io])

[blog.alexellis.io]: https://blog.alexellis.io/golang-writing-unit-tests/

## "Testify" package, Mat Ryer

The [assert] package provides some helpful methods that allow you to write better test code in Go.

[assert]: https://github.com/stretchr/testify

- Prints friendly, easy to read failure descriptions
- Allows for very readable code
- Optionally annotate each assertion with a message

```
// testify/assert
package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func IsSuperAnimal(animal string) bool {
	return strings.ToLower(animal) == "gopher"
}

func TestIsSuperAnimal(t *testing.T) {
	assert.True(t, IsSuperAnimal("gopher"))
}
```
- [Example from blog.friendsofgo.tech](https://blog.friendsofgo.tech/posts/empezando-con-los-tests-automatizados-en-go/)

## "Is" package, Mat Ryer

Package [is] provides a lightweight extension to the standard library's testing capabilities.

Comments on the assertion lines are used to add a description.

[is]: https://github.com/matryer/is

```
// is
package main

import (
	"github.com/matryer/is"
	"strings"
	"testing"
)

func IsSuperAnimal(animal string) bool {
	return strings.ToLower(animal) == "gopher"
}

func TestIsSuperAnimal(t *testing.T) {
	is := is.New(t)
	is.True(IsSuperAnimal("gopher"))
}

```

With `is.NewRelaxed(t)` instead of `is.New(t)`, relaxed mode, failures call `T.Fail` allowing multiple failures per test.

- [Example from blog.friendsofgo.tech](https://blog.friendsofgo.tech/posts/empezando-con-los-tests-automatizados-en-go/)
- [Other examples](../src/13-testing/03-testing-packages)

## Mocks
*We often need to simulate or mimic an object to create a deterministic, fast and network-independent object. Such an object is useful while testing. This practice is also known as mocking.*
- [How to mock in go, Adnaan Badr](http://adnaan.badr.in/blog/2019/01/10/how-to-mock-in-go/)
- [Mocking dependencies in go](https://medium.com/agrea-technogies/mocking-dependencies-in-go-bb9739fef008)
- https://github.com/golang/mock
- https://github.com/matryer/moq
- https://github.com/DATA-DOG/go-sqlmock

## Benchmarking

<small>From: Dave Cheney</small>

Given a method:
```
func Fib(n int) int {
        if n < 2 {
                return n
        }
        return Fib(n-1) + Fib(n-2)
}
```

*Benchmarks are placed inside `_test.go files and follow the rules of their Test counterparts. 
In this first example it is going to benchmark the speed of computing the 10th number in the Fibonacci series.*

```
// from fib_test.go
func BenchmarkFib10(b *testing.B) {
        // run the Fib function b.N times
        for n := 0; n < b.N; n++ {
                Fib(10)
        }
}
```
Launch a benchmark:

```bash
go test -bench=.
```

Response

```
goos: darwin
goarch: amd64
pkg: github.com/slaparra/go-training/src/13-testing/04-benchmark/somePackage
BenchmarkPrintAString-4   	10000000	       152 ns/op
PASS
ok  	github.com/slaparra/go-training/src/13-testing/04-benchmark/somePackage	1.692s
```
Writing a benchmark is very similar to writing a test as they share the infrastructure from the testing package. Some of the key differences are

- Benchmark functions start with **Benchmark** not Test.
- Benchmark functions are run several times by the testing package. The value of b.N will increase each time until the benchmark runner is satisfied with the stability of the benchmark.
- Each benchmark must execute the code under test b.N times. The for loop in BenchmarkFib10 will be present in every benchmark function.

#### Links:
- [How to write benchmarks in Go, Dave Cheney](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [Analyzing the performance of Go functions with benchmarks, Francesc Campoy](https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6)

---

## Videos

- [**An Introduction to Testing in Go**](https://www.youtube.com/watch?v=GlA57dHa5Rg)
- [Tests & benchmark, Learn to code](https://www.youtube.com/watch?v=OtNJgykLMjk)
- [Unit testing HTTP servers, Francesc campoy](https://www.youtube.com/watch?v=hVFEV-ieeew)
- [The Art of Testing, Mat Ryer](https://www.youtube.com/watch?v=EOpj9aZ8Kfo)
- [Quality Tests in Go, Mat Ryer](https://www.youtube.com/watch?v=MMnaq2jwAiE)
- [Advanced Testing with Go, Mitchell Hashimoto](https://www.youtube.com/watch?v=yszygk1cpEc)
- [Unit Testing in Go, advance examples by packtpub.com](https://www.youtube.com/watch?v=j_Isq09hZG8)
- [Basic Benchmarks, justforfunc](https://www.youtube.com/watch?time_continue=97&v=2AulMm-hsdI)

## Links

- [Testing, Golang book](http://www.golang-book.com/books/intro/12) 
- [Package testing, Godoc](https://godoc.org/testing)
- [Testable Examples in Go, Golang blog](https://blog.golang.org/examples)
- [How to test with Go](https://www.calhoun.io/how-to-test-with-go/)
- [Writing unit tests, Alex Ellis](https://blog.alexellis.io/golang-writing-unit-tests/)
- [Starting with automated tests, Friends of Go](https://blog.friendsofgo.tech/posts/empezando-con-los-tests-automatizados-en-go/) (spanish)
