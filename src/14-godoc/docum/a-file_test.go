package docum

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	value := Start()
	expected := "Santi"

	if value != expected {
		t.Errorf("Got %s, expected %s", value, expected)
	}
}

// This method will put an example in the documentation
func ExampleStart() {
	value := Start()
	fmt.Println(value)
	// Output:
	// Santi (message written in test file with prefix Example<Method>)
}
