# Testing

Go includes a special program that makes writing tests easier.
- Test file convention: `fileToTest_test.go`  
- The test file should be in the same package as the one being tested  
- Be in a function `func TestXXX(t *testing.T)`.

Run tests:
`go test` or `go test -v` to get more info

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

- https://golang.org/pkg/testing/

### Table driven tests
to do

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

- https://www.golang-book.com/books/intro/12
- https://blog.golang.org/examples
- https://www.calhoun.io/how-to-test-with-go/
- https://blog.alexellis.io/golang-writing-unit-tests/
- https://blog.friendsofgo.tech/posts/empezando-con-los-tests-automatizados-en-go/ (spanish)
