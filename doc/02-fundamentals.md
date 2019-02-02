# Fundamentals

- [Writing, building, installing, and testing Go code (video)](https://www.youtube.com/watch?v=XCsL89YtqCs)
- [Code examples](../src/01-fundamentals): array-slice, constants, print, variables...
- [How do you structure your go apps (video)](https://www.youtube.com/watch?v=B5oQnECDJ8g)
- [Structure & maintain a Golang project with multiple files, stackoverflow](https://stackoverflow.com/questions/48276724/structure-maintain-a-golang-project-with-multiple-files)

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
    - [Constants](https://golang.org/ref/spec#Constant_declarations)
    - [Numeric types](https://golang.org/ref/spec#Numeric_types)
    - [Package types](https://golang.org/pkg/go/types/)  
    
#### Simple quote
```
e := "Hello"
f := `Do you like my hat?`
g := 'M'

fmt.Printf("%T \n", e)
fmt.Printf("%T \n", f)
fmt.Printf("%T \n", g)

a := `this is a
string literal
and you can print "quotes"
like this`

(string)
(string)
(int32)!!!!
```
#### Arithmetic operations

```
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers

&    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&^   bit clear (AND NOT)    integers

<<   left shift             integer << unsigned integer
>>   right shift            integer >> unsigned integer
```

https://golang.org/ref/spec#Arithmetic_operators

**Package math** provides basic constants and mathematical functions: https://golang.org/pkg/math

#### Constants
```
const Pi float64 = 3.14159265358979323846
const zero = 0.0         // untyped floating-point constant
const (
	size int64 = 1024
	eof        = -1  // untyped integer constant
)
const a, b, c = 3, 4, "foo"  // a = 3, b = 4, c = "foo", untyped integer and string constants
const u, v float32 = 0, 3    // u = 0.0, v = 3.0
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

#### Shift 1 bit to the left

```
a := 42
b := a << 1
fmt.Printf("%d", b) //84
fmt.Printf("%b", b) //1010100
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

![Picture](../resources/scopes.1359063689.png)  
<sub>(Image: [golang-book.com](https://www.golang-book.com/books/web/01-02))</sub>

Things defined in the function scope have access to anything above them (file, package, universe), but the reverse is not true. A variable defined in a function is only accessible within that function or blocks defined inside of it.

## Types

A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a type name, if it has one, or specified using a type literal, which composes a type from existing types.

Go is strongly typed programming language. It’s strict when it goes to types and mistakes will be reported during compilation

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

- [Golang Basic Types, Operators](https://www.callicoder.com/golang-basic-types-operators-type-conversion/#type-conversion)

### Type conversion

*Unlike in C, in Go assignment between items of different type requires an explicit conversion*
- [Example type custom struct conversion](../src/01-fundamentals/type-conversion.go)
- [Type conversions in Tour of go](https://tour.golang.org/basics/13)

*There are times where it might be desirable to convert value to other type. Golang doesn’t allow to do it in arbitrary way. They’re certain rules enforced by the type system.*

- [Conversions in go, Golang spec](https://medium.com/golangspec/conversions-in-go-4301e8d84067)
- [Type conversions](https://www.callicoder.com/golang-basic-types-operators-type-conversion/#type-conversion)

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
