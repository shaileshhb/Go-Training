package maths

import "testing"

func TestAdd(t *testing.T) {

	actual := add(20, 30)
	var expected int64 = 50

	if actual != expected {
		t.Error("Acutal:", actual, "Bt expected is: ", expected)
	}
}

func TestAddValuesInBulk(t *testing.T) {

	var list = []struct {
		firstNum  int
		secondNum int
		expected  int64
	}{
		{
			10, 20, 30,
		},
		{
			-10, 20, 10,
		},
		{
			10, -20, -10,
		},
		{
			10, 40, 50,
		},
		{
			-10, -20, -30,
		},
	}

	for _, num := range list {
		if actual := add(num.firstNum, num.secondNum); actual != num.expected {
			t.Error("Acutal:", actual, "Bt expected is: ", num.expected)
		}
	}
}
