# Control structures

## Conditional
- [if examples](../src/05-control-structures/conditional.go)
- [Golang spec: If statements](https://golang.org/ref/spec#If_statements)
- [Effective go: If](https://golang.org/doc/effective_go.html#if)

```
if x%2 == 1 {
    fmt.Printf("X (%d) is odd\n", x)
} else {
    fmt.Printf("X (%d) is even\n", x)
}
```

## Loop
https://golang.org/doc/effective_go.html#for  
https://golang.org/ref/spec#For_statements  
[Loop examples](../src/05-control-structures/loop.go)
```
    for i := 32; i < 123; i++ {
        fmt.Printf("%d \t %b \t %x \t %q\n", i, i, i, i) //%q ascii
    }
```

## Range
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

## Switch
*Go `Switch` only runs the selected case, not all the cases that follow. The break statement is provided automatically in Go.*
```go
func main() {    
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.", os)
    }
}
```

To execute a code for different cases:
```go
switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}

switch x := f(); {  // missing switch expression means "true"
case x < 0: return -x
default: return x
}

switch {
case x < y: f1()
case x < z: f2()
case x == 4: f3()
}

```
- [A tour of go](https://tour.golang.org/flowcontrol/9)
- [Switch examples](../src/05-control-structures/switch.go)
- [Doc Switch statements](https://golang.org/ref/spec#Switch_statements)
- [Effective go: Switch](https://golang.org/doc/effective_go.html#switch)