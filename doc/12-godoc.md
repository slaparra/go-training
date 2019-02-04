# Go documentation

## Go doc command

`go doc` prints the documentation for a package, const, func, type, var, or method (high level view)

`go doc` accepts zero, one, or two arguments.

#### Zero:
```
go doc
```

Show the documentation for the package in the current directory.

#### One argument:

```
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
```
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
```
godoc fmt

# to see the source
godoc -src fmt Printf
```

## Export documentation to godoc.org
Paste the package github url to [https://godoc.org] and automatically will appear in it   

[https://godoc.org]: https://godoc.org

---

# Links
- https://blog.golang.org/godoc-documenting-go-code
- https://golang.org/cmd/doc/
