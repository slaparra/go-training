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

## Switch
*Go `Switch` only runs the selected case, not all the cases that follow. The break statement  is provided automatically in Go.*
```
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
- [A tour of go](https://tour.golang.org/flowcontrol/9)
- [Switch examples](../src/05-control-structures/switch.go)
- [Doc Switch statements](https://golang.org/ref/spec#Switch_statements)
- [Effective go: Switch](https://golang.org/doc/effective_go.html#switch)