package executor

import (
	"fmt"
	"reflect"
	"strings"
)

type carDetails struct {
	regNum, color string
}

var slots []carDetails

// CreateParkingSlots ...
func CreateParkingSlots(slotNum int) {
	slots = make([]carDetails, slotNum)
}

// ParkCar ...
func ParkCar(regNum, color string) {
	slotsNum := len(slots)
	slotNo := slotsNum
	newCar := carDetails{
		regNum: regNum,
		color:  color,
	}

	for slot, detail := range slots {
		if detail.regNum == newCar.regNum {
			fmt.Printf("Already in slot no %d\n", slot+1)
			return
		}
		if reflect.DeepEqual(detail, carDetails{}) {
			slotNo = min(slot, slotNo)
		}
	}

	if slotNo >= slotsNum {
		fmt.Println("No slot available :(")
		return
	}
	slots[slotNo] = newCar
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// LeaveParking ...
func LeaveParking(slotNum int) {
	if slotNum > len(slots) || slotNum < 0 {
		fmt.Println("Invalid slot no :(")
		return
	}
	if reflect.DeepEqual(slots[slotNum-1], carDetails{}) {
		fmt.Printf("No vehicle found in slot %d\n", slotNum)
		return
	}

	slots[slotNum-1] = carDetails{}
}

// PrintStatus ...
func PrintStatus() {
	fmt.Println("Slot\tRegistration Number\tColor")
	for slot, details := range slots {
		if details != (carDetails{}) {
			fmt.Printf("%-6v\t%-20s\t%-16s\n", slot+1, details.regNum, details.color)
		}
	}
}

// FindRegNumsByColor ...
func FindRegNumsByColor(color string) {
	for _, details := range slots {
		if strings.EqualFold(details.color, color) {
			fmt.Printf("%s\n", details.regNum)
		}
	}
}

// FindSlotNumsByColor ...
func FindSlotNumsByColor(color string) {
	for slot, details := range slots {
		if strings.EqualFold(details.color, color) {
			fmt.Printf("%d\n", slot+1)
		}
	}
}

// FindSlotNoByRegNum ...
func FindSlotNoByRegNum(regNum string) {
	for slot, details := range slots {
		if strings.EqualFold(details.regNum, regNum) {
			fmt.Printf("%d\n", slot+1)
		}
	}
}
