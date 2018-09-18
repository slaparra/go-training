# Control structures

## Conditional
- [if examples](../src/05-control-structures/conditional.go)
- [Golang spec: If statements](https://golang.org/ref/spec#If_statements)
- [Effective go: If](https://golang.org/doc/effective_go.html#if)

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
- [Switch examples](../src/05-control-structures/switch.go)
- [Doc Switch statements](https://golang.org/ref/spec#Switch_statements)
- [Effective go: Switch](https://golang.org/doc/effective_go.html#switch)