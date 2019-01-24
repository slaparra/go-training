package factorial

import "testing"

func TestFactorialDataDrivenTest(t *testing.T) {
	type dataSource struct {
		numberToCheck  int
		expectedResult int
	}

	dataSourceTable := []dataSource{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
	}

	for _, x := range dataSourceTable {
		value := x.numberToCheck

		if factorial(value) != x.expectedResult {
			t.Errorf("Expected %d, Got %d", x.expectedResult, value)
		}
	}
}
