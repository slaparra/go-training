// is
package tests

import (
	"errors"
	"github.com/matryer/is"
	"strings"
	"testing"
)

func isSignedIn(aVar string) (bool, error) {
	result := strings.ToLower(aVar) == "a String"

	if !result {
		return false, errors.New("force an error")
	}

	return result, nil
}

func readBody(body string) string {
	return body
}

func Test1(t *testing.T) {
	isVar := is.New(t)

	_, err := isSignedIn("an int")
	isVar.NoErr(err) // isSignedIn error
}

func Test2(t *testing.T) {
	isVar := is.New(t)

	signedIn, _ := isSignedIn("a String")
	isVar.Equal(signedIn, true) // must be signed in
}

func Test3(t *testing.T) {
	isVar := is.New(t)

	body := readBody("Hi no there")
	isVar.True(strings.Contains(body, "Hi there")) //an error message
}
