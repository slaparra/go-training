package author

import "fmt"

// fields are not visible (first letter is lowercase) so we can't access from main.go
// only go files in the same package can access to the not visible fields
type Author struct {
	firstName string
	lastName  string
	bio       string
}

func (a *Author) FullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

func (a *Author) Bio() string {
	return fmt.Sprintf("Bio: %s_%s: %s", a.firstName, a.lastName, a.bio)
}

func (a *Author) ChangeName(aName string) {
	a.firstName = aName
}

// we need to create a "named constructor" to create an Author object outside the package
// method will be Visible outside the package
func Create(aFirstName string, aLastName string, aBio string) *Author {
	return &Author{aFirstName, aLastName, aBio}
}
