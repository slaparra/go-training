# Fundamentals

- [Writing, building, installing, and testing Go code (video)](https://www.youtube.com/watch?v=XCsL89YtqCs)
- [Code examples](../src/01-fundamentals): array-slice, constants, print, variables...
- [How do you structure your go apps](https://www.youtube.com/watch?v=B5oQnECDJ8g)
## Format
https://godoc.org/fmt
```
fmt.Printf("%d - %b\n", 43, 17) #decimal - binary
fmt.Printf("%s\n", "github.com/slaparra") #string
```

## Variables

- Default assignment to a var is zero value. Each type has its own zero value
    - [Golang tour](https://tour.golang.org/basics/12)
    - [Go spec](https://golang.org/ref/spec#The_zero_value)
- All variables **are always passed by value** to another functions
    - [Yury Pitsishin post](http://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/)
    - [Dave Cheney post](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go)
- [Golang: Variables](https://golang.org/ref/spec#Variables)
- [Example: How to declare variables](../src/01-fundamentals/variables2.go)
- Types [(spec)](https://golang.org/ref/spec#Types)
    - [Strings](https://blog.golang.org/strings)
    - [Array in data structures section](06-data-structures.md)       
    - [Slice in data structures section](06-data-structures.md)
    - [Numeric types](https://golang.org/ref/spec#Numeric_types)
    - [Package types](https://golang.org/pkg/go/types/)  
    
Simple quote:
```
e := "Hello"
f := `Do you like my hat?`
g := 'M'

fmt.Printf("%T \n", e)
fmt.Printf("%T \n", f)
fmt.Printf("%T \n", g)

(string)
(string)
(int32)!!!!
```

#### Print variables
https://godoc.org/fmt  

```
fmt.Printf("%v",aVar)  
The default format for %v is:

bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
``` 

## Scope

References:
- [Scope golang-book, with great examples**](https://www.golang-book.com/books/web/01-02#scope)
- [Golang: Declarations and scope](https://golang.org/ref/spec#Declarations_and_scope)
- [Todd McLeod scope & closure examples](../src/03-block-scope)

Definition:
- The scope of a predeclared identifier is the universe block.
- The scope of an identifier denoting a constant, type, variable, or function (but not method) declared at top level (outside any function) is the package block.
- The scope of the package name of an imported package is the file block of the file containing the import declaration.
- The scope of an identifier denoting a method receiver, function parameter, or result variable is the function body.
- The scope of a constant or variable identifier declared inside a function begins at the end of the ConstSpec or VarSpec (ShortVarDecl for short variable declarations) and ends at the end of the innermost containing block.
- The scope of a type identifier declared inside a function begins at the identifier in the TypeSpec and ends at the end of the innermost containing block.

![Picture](https://www.golang-book.com/public/img/web/scopes.0.png)  

Things defined in the function scope have access to anything above them (file, package, universe), but the reverse is not true. A variable defined in a function is only accessible within that function or blocks defined inside of it.

## Types

A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a type name, if it has one, or specified using a type literal, which composes a type from existing types.

```
type (
	B1 string
	B2 B1
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
}

aPerson := Employee{"John", "Raimon", 28}
```
## Blank identifier
Basically a blank identifier is like a box where you can put things if you don't need them.  

It's a bit like writing to the Unix /dev/null file: it represents a write-only value to be used as a place-holder where a variable is needed but the actual value is irrelevant (in code is the underscore).

[More documentation and uses of Blank identifier](https://golang.org/doc/effective_go.html#blank)

```
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//blank identifier is the underscore _
func main() {
	res, _ := http.Get("http://www.geekwiseacademy.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
```

## Constants, Iota
- [Constants: at Golang blog](https://blog.golang.org/constants)
- [Golang Constants](https://golang.org/ref/spec#Constant_declarations)  
- [Iota](https://golang.org/ref/spec#Iota) 
- [Constants & Iota example](../src/01-fundamentals/constants.go)

## Rune literals
https://golang.org/ref/spec#Rune_literals
What's a rune: 
- Is Go terminology for a single Unicode code point
- An integer value identifying a Unicode code point
- It's a character
- Alias for int32 ([Numeric types](https://golang.org/ref/spec#Numeric_types))
    - 4 bytes -> 4 * 8 = 32
    - UTF-8 is a 4 byte coding scheme
    - with int 32 we have numbers for each of the code points
    
** from Todd McLeod
    
A rune literal is expressed as one or more characters enclosed in single quotes, as in 'x' or '\n'.

```
var aRuneVar rune = 'h'
```

```
fmt.Println([]byte("Hello"))

//[72 101 108 108 111] 1 byte per character
```
https://en.wikipedia.org/wiki/ASCII#Printable_characters

```
fmt.Println([]byte("世界"))

//[228 184 150 231 149 140] 3 bytes per character
```

```
word := "hello"
letter := rune(word[0])
fmt.Println(letter)      //72
```
[Rune & UTF-8 example](../src/01-fundamentals/rune.go)
