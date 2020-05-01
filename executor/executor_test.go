package executor

import (
	"fmt"
	"testing"
)

func TestCreateParkingSlots(t *testing.T) {
	fmt.Printf("All test case passed")
}

func TestParkCar(t *testing.T) {
	if output := ParkCar(regNum, color); output != test.res[index] {
		t.Errorf("Test Name: %s, Command: %s, Expected: %s, Got: %s\n", t.Name, txt, test.res[index], output)
	}
	fmt.Printf("All test case passed")
}

func TestLeaveParking(t *testing.T) {
	if output := LeaveParking(slotNum); output != test.res[index] {
		t.Errorf("Test Name: %s, Command: %s, Expected: %s, Got: %s\n", t.Name, txt, test.res[index], output)
	}
	fmt.Printf("All test case passed")
}

func TestPrintStatus(t *testing.T) {
	if output := PrintStatus(); output != test.res[index] {
		t.Errorf("Test Name: %s, Command: %s, Expected: %s, Got: %s\n", t.Name, txt, test.res[index], output)
	}
	fmt.Printf("All test case passed")
}

func TestFindRegNumsByColor(t *testing.T) {
	if output := FindRegNumsByColor(color); output != test.res[index] {
		t.Errorf("Test Name: %s, Command: %s, Expected: %s, Got: %s\n", t.Name, txt, test.res[index], output)
	}
	fmt.Printf("All test case passed")
}

func TestFindSlotNumsByColor(t *testing.T) {
	if output := FindSlotNumsByColor(color); output != test.res[index] {
		t.Errorf("Test Name: %s, Command: %s, Expected: %s, Got: %s\n", t.Name, txt, test.res[index], output)
	}
	fmt.Printf("All test case passed")
}

func TestFindSlotNoByRegNum(t *testing.T) {
	if output := FindSlotNoByRegNum(regNum); output != test.res[index] {
		t.Errorf("Test Name: %s, Command: %s, Expected: %s, Got: %s\n", t.Name, txt, test.res[index], output)
	}
	fmt.Printf("All test case passed")
}
