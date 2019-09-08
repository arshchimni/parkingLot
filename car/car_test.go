package car

import (
	"testing"
)

func TestCar(t *testing.T) {
	car := NewCar("testCar", "white")

	expected := "testCar"
	actual := car.GetRegNo()
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	expected = "white"
	actual = car.GetColor()
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

}
