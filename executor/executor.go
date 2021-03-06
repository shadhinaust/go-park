package executor

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const createSlot = "create_parking_lot"
const park = "park"
const leaveSlot = "leave"
const status = "status"
const regNumsForColor = "registration_numbers_for_cars_with_colour"
const slotNumsForColor = "slot_numbers_for_cars_with_colour"
const slotNumForRegNum = "slot_number_for_registration_number"

type carDetails struct {
	regNum, color string
}

var slots []carDetails

// ExecCmd ...
func ExecCmd(cmd []string) {
	switch cmd[0] {
	case createSlot:
		slotNum, _ := strconv.Atoi(cmd[1])
		CreateParkingSlots(slotNum)
	case park:
		regNum := cmd[1]
		color := cmd[2]
		if res := ParkCar(regNum, color); res != "" {
			fmt.Println(res)
		}
	case leaveSlot:
		slotNum, _ := strconv.Atoi(cmd[1])
		if res := LeaveParking(slotNum); res != "" {
			fmt.Println(res)
		}
	case status:
		fmt.Println("Slot\tRegistration Number\tColor")
		if res := PrintStatus(); res != "" {
			fmt.Println(res)
		}
	case regNumsForColor:
		color := cmd[1]
		fmt.Println(FindRegNumsByColor(color))
	case slotNumsForColor:
		color := cmd[1]
		fmt.Println(FindSlotNumsByColor(color))
	case slotNumForRegNum:
		regNum := cmd[1]
		fmt.Println(FindSlotNoByRegNum(regNum))
	}
}

// CreateParkingSlots ...
func CreateParkingSlots(slotNum int) {
	slots = make([]carDetails, slotNum)
}

// ParkCar ...
func ParkCar(regNum, color string) string {
	slotsNum := len(slots)
	slotNo := slotsNum
	newCar := carDetails{
		regNum: regNum,
		color:  color,
	}

	for slot, detail := range slots {
		if detail.regNum == newCar.regNum {
			return fmt.Sprintf("Already in slot no %d", slot+1)
		}
		if reflect.DeepEqual(detail, carDetails{}) {
			slotNo = min(slot, slotNo)
		}
	}

	if slotNo >= slotsNum {
		return "No slot available :("
	}
	slots[slotNo] = newCar
	return ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// LeaveParking ...
func LeaveParking(slotNum int) string {
	if slotNum > len(slots) || slotNum < 0 {
		return "Invalid slot no :("
	}
	if reflect.DeepEqual(slots[slotNum-1], carDetails{}) {
		return fmt.Sprintf("No vehicle found in slot %d", slotNum)
	}
	slots[slotNum-1] = carDetails{}
	return ""
}

// PrintStatus ...
func PrintStatus() string {
	status := ""
	fmt.Println("Slot\tRegistration Number\tColor")
	for slot, details := range slots {
		if details != (carDetails{}) {
			status += fmt.Sprintf("%-6v\t%-20s\t%-16s\n", slot+1, details.regNum, details.color)
		}
	}
	return status
}

// FindRegNumsByColor ...
func FindRegNumsByColor(color string) string {
	for _, details := range slots {
		if strings.EqualFold(details.color, color) {
			return details.regNum
		}
	}
	return "Not Found :("
}

// FindSlotNumsByColor ...
func FindSlotNumsByColor(color string) string {
	for slot, details := range slots {
		if strings.EqualFold(details.color, color) {
			return strconv.Itoa(slot + 1)
		}
	}
	return "Not Found :("
}

// FindSlotNoByRegNum ...
func FindSlotNoByRegNum(regNum string) string {
	for slot, details := range slots {
		if strings.EqualFold(details.regNum, regNum) {
			return strconv.Itoa(slot + 1)
		}
	}
	return "Not Found :("
}
