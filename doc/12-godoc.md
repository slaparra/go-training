# Go documentation

## Go doc command

`go doc` prints the documentation for a package, const, func, type, var, or method (high level view)

`go doc` accepts zero, one, or two arguments.

#### Zero:

```bash
go doc
```

Show the documentation for the package in the current directory.

#### One argument:

```bash
go doc <pkg>
go doc <sym>[.<methodOrField>]
go doc [<pkg>.]<sym>[.<methodOrField>]
go doc [<pkg>.][<sym>.]<methodOrField>

# example 
go doc fmt
go doc fmt.Errorf

# more examples in
go help doc 
```
The first item in this list that succeeds is the one whose documentation is printed. If there is a symbol but no package, the package in the current directory is chosen. However, if the argument begins with a capital letter it is always assumed to be a symbol in the current directory.

#### Two arguments:
```bash
go doc <pkg> <sym>[.<methodOrField>]
```
Show the documentation for the package, symbol, and method or field. The first argument must be a full package path. This is similar to the command-line usage for the godoc command.

- [Example main.go and package docum](../src/14-godoc/main.go)


## Godoc command

Godoc **extracts and generates documentation** for Go programs.

It runs as a web server and presents the documentation as a web page.

`godoc -http=:6060`

Creates documentation for your own programs as well

https://godoc.org/golang.org/x/tools/cmd/godoc

The `godoc <package>` command shows documentation for the package more complete 

```bash
godoc fmt

# to see the source
godoc -src fmt Printf
```

#### How to document 
- A type, variable, constant, function, or package, 
- Write a comment directly preceding its declaration, with no intervening blank line.
    - Begin with the name of the element
    - For packages: 
        - First sentence appears in package list
        - If a large amount of documentation, place in its own file doc.go

##### Create examples in doc
Write in the test file a method with Example prefix:

```go
// This method will put an example in the documentation
func ExampleStart() {
    value := Start()
    fmt.Println(value)
    // Output:
    // The expected output (message written in test file with prefix Example<Method>)
}
```

Will put an example in the doc:

```
▾ Example

This method will put an example in the documentation

Code:

value := Start()
fmt.Println(value)

Output:

The expected output (message written in test file with prefix Example<Method>)
```

## Export documentation to godoc.org
Paste the package github url to [https://godoc.org] and automatically will appear in it   

[https://godoc.org]: https://godoc.org

---

# Links
- https://blog.golang.org/godoc-documenting-go-code
- https://golang.org/cmd/doc/
