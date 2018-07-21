
## Format
https://godoc.org/fmt
```
fmt.Printf("%d - %b\n", 43, 17) #decimal - binary
fmt.Printf("%s\n", "github.com/slaparra") #string
```

## Loop
https://golang.org/ref/spec#For_statements
https://golang.org/doc/effective_go.html#for
```
	for i := 32; i < 123; i++ {
		fmt.Printf("%d \t %b \t %x \t %q\n", i, i, i, i) //%q ascii
	}
```

## Packages
- [Packages (book An Introduction to Programming in Go )](https://www.golang-book.com/books/intro/11)
- [Package file examples](../todd-mcleod/02-package)
- Methods and vars have to be capitalized to be visible outside the package
