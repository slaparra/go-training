# OOP in Go

- [Composition Instead of Inheritance - OOP in Go](https://golangbot.com/inheritance/)
- [Composition with go](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html)

In go we don't have public/protected/private visibility in methods nor fields.  
Only fields or methods can be visible or not visible outside the package (with capital letter at the first).  

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

or we can create a named constructor:
```
func Create(aFirstName string, aLastName string, aBio string) Author {
	return Author{aFirstName, aLastName, aBio}
}
```