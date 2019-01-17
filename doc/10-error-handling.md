# Error handling

*In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C.*

- [GoBlog: Error handling and go](https://blog.golang.org/error-handling-and-go)
- [GoByExample: Errors](https://gobyexample.com/errors)

```
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f
```

There are 4 forms of handling errors:
```
_, err := os.Open("no-file.txt")	//try to find a non existing txt file to force an error
if err != nil {
    // fmt.Println("err happened", err)
    // log.Println("err happened", err)
    // log.Fatalln(err)
    // panic(err)
}
```
*`log.Println` calls Output to print to the standard logger. Arguments are handled in the manner of `fmt.Println`.*

*`Fatalln` is equivalent to `log.Println()` followed by a call to `os.Exit(1)`.*  [Exit into](https://godoc.org/os#Exit)

#### Panic

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

- [GoByExample: Panic](https://gobyexample.com/panic)
- [Defer, panic & recover](https://www.golang-book.com/books/intro/7#section6)

Video
- [Dont Just Check Errors Handle Them Gracefully, Dave Cheney](https://www.youtube.com/watch?v=lsBF58Q-DnY)