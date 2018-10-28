package author

import "fmt"

// Author is a struct to contain the mvp data
// fields are not visible (first letter is lowercase) so we can't access from main.go
// only go files in the same package can access to the not visible fields
type Author struct {
	firstName string
	lastName  string
	bio       string
}

// FullName method prints the name plus lastName
func (a *Author) FullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

// Bio method prints fullName plus bio
func (a *Author) Bio() string {
	return fmt.Sprintf("Bio: %s_%s: %s", a.firstName, a.lastName, a.bio)
}

// ChangeName method changes the name with the firstName
func (a *Author) ChangeName(aName string) {
	a.firstName = aName
}

// Create method returns the address of a new Author
// we need to create a "named constructor" to create an Author object outside the package
// method will be Visible outside the package
func Create(aFirstName string, aLastName string, aBio string) *Author {
	return &Author{aFirstName, aLastName, aBio}
}
