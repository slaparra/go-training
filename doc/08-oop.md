# OOP in Go

- [Composition Instead of Inheritance - OOP in Go](https://golangbot.com/inheritance/)
- [Composition with go](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html)
- [Notes about opp in go by Todd McLeod](https://github.com/GoesToEleven/GolangTraining/blob/master/20_struct/00_object-oriented/notes.txt)

## Structs

*A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-blank field names must be unique.*

- **[Golangbot, structs](https://golangbot.com/structs/)**
- [Golang book, structs](https://www.golang-book.com/books/intro/9#section1)
- [Golang reference spec](https://golang.org/ref/spec#Struct_types)

```
// named structure
type author struct {  
    firstName string
    lastName  string
    bio       string
}

func (a author) fullName() string {  
    return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {  
    title   string
    content string
    author
}
```

It is possible to declare structures without declaring a new type and these type of structures are called **anonymous structures**

```
var employee struct {  
        firstName, lastName string
        age int
}
```

## Objects

In go we don't have public/protected/private visibility in methods nor fields.  
Only fields or methods can be exported or not-exported (visible or not visible) outside the package (with capital letter at the first).  

See [author example](../todd-mcleod/09-oop/01-encapsulate_object)

```
package main

import (  
    "fmt"
)

// create the object with fields not visible
type author struct {  
    firstName string
    lastName  string
    bio       string
}

// getter method Visible
func (a author) FullName() string {  
    return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}
```

How to instantiate an object with go:

```
// if author is in the same package
aVar := Author{"aFirstName", "aLastName", "aBio"}
```

or we can create a named constructor (like a static method in other languages):
```
func Create(aFirstName string, aLastName string, aBio string) Author {
	return Author{aFirstName, aLastName, aBio}
}
```

## Inheritance in go

Extracted from:
- [Struct, promoted fields](https://golangbot.com/structs/)
- [Composition Instead of Inheritance - OOP in Go](https://golangbot.com/inheritance/)

*Go does not support inheritance, however it does support composition. The generic definition of composition is "put together". One example of composition is a car. A car is composed of wheels, engine and various other parts.*

*Fields that belong to a anonymous struct field in a structure are called **promoted** fields since they can be accessed as if they belong to the structure which holds the anonymous struct field. I can understand that this definition is quite complex so lets dive right into some code to understand this :).*

```
type Address struct {  
    city, state string
}

type Person struct {  
    name string
    age  int
    Address
}
```

*In the above code snippet, the Person struct has an anonymous field Address which is a struct. Now the fields of the Address struct namely city and state are called promoted fields since they can be accessed as if they are directly declared in the Person struct itself.*
```
package main

import (  
    "fmt"
)

type Address struct {  
    city, state string
}
type Person struct {  
    name string
    age  int
    Address
}

func main() {  
    var p Person
    p.name = "Naveen"
    p.age = 50
    p.Address = Address{
        city:  "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.city) //city is promoted field
    fmt.Println("State:", p.state) //state is promoted field
}

```

- [Promoted code example file](../todd-mcleod/09-oop/02-promoted-inheritance/)