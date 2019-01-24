# Testing

Go includes a special program that makes writing tests easier.
- Test file convention: `fileToTest_test.go`  
- The test file should be in the same package as the one being tested  
- Be in a function `func TestXXX(t *testing.T)`.
- Use `t.Error` for signal failure

Run tests:
`go test` or `go test -v` to run tests with verbose mode

### Package testing

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
Error is equivalent to Log followed by Fail.
- https://golang.org/pkg/testing/
- https://godoc.org/testing

### Table driven tests
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

*At the heart of all table driven tests is the table itself, which provides the inputs and expected results of the function under test. In most cases the table is a slice of anonymous structs, which allows the table to be written in a compact form:*
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

### Testify package, Mat Ryer
to do

### Is package, Mat Ryer
to do


## Videos

- [Unit testing HTTP servers, Francesc campoy](https://www.youtube.com/watch?v=hVFEV-ieeew)
- [The Art of Testing, Mat Ryer](https://www.youtube.com/watch?v=EOpj9aZ8Kfo)
- [Quality Tests in Go, Mat Ryer](https://www.youtube.com/watch?v=MMnaq2jwAiE)
- [Advanced Testing with Go, Mitchell Hashimoto](https://www.youtube.com/watch?v=yszygk1cpEc)

## Links

- https://godoc.org/testing
- https://www.golang-book.com/books/intro/12
- https://blog.golang.org/examples
- https://www.calhoun.io/how-to-test-with-go/
- https://blog.alexellis.io/golang-writing-unit-tests/
- https://blog.friendsofgo.tech/posts/empezando-con-los-tests-automatizados-en-go/ (spanish)
