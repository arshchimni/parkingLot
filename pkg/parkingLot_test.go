package pkg

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestLeaveWithError(t *testing.T) {
	expected := "Parking lot not constructed yet cannot park"
	actual := Leave(1).Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestParkWithError(t *testing.T) {
	expected := "Parking lot not constructed yet cannot park"
	actual := Park("a", "red").Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestCreateParkingLotError(t *testing.T) {
	expected := "cannot create parking lot with given slots"
	actual := CreateParkingLot(0).Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestCreateParkingLotoutError(t *testing.T) {
	expected := ""
	actual := CreateParkingLot(2)
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual.Error())
	}
}

func TestParkWithoutError(t *testing.T) {
	expected := ""
	actual := Park("b", "red")
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual.Error())
	}
}

func TestSlotNoForRegNo(t *testing.T) {
	expected := ""
	actual := SlotNumberForRegistrationNumber("b")
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%v'", expected, actual)
	}
}

func TestRegNosForCarsWithColor(t *testing.T) {
	expected := ""
	actual := RegistrationNumbersForCarsWithColour("red")

	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%v'", expected, actual)
	}

}

func TestRegNosForCarsWithColor2(t *testing.T) {
	expected := "Car with that color not present in the parking lot"
	err := RegistrationNumbersForCarsWithColour("white")

	if err != nil && err.Error() != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, err.Error())
	}

}

func TestStatus(t *testing.T) {
	// Try to Pipe the stdout of the program to match the status output against what is expected
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	err := Status()
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	expected := "Slot No.\tRegistration No.\tColour\n1\t\tB\t\tred\n"
	if err != nil || expected != out {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, out)
	}
}

func TestSlotNosForCarsWithColor(t *testing.T) {
	expected := ""
	actual := SlotNumbersForCarsWithColour("red")

	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%v'", expected, actual)
	}

}

func TestLeaveWithoutError(t *testing.T) {
	expected := ""
	actual := Leave(1)
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual.Error())
	}
}

func TestSlotNumberForRegistrationNumberNotFound(t *testing.T) {

	expected := "Not found"
	actual := SlotNumberForRegistrationNumber("b")
	if actual.Error() != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%v'", expected, actual)
	}
}
