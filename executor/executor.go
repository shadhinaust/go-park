package executor

import (
	"fmt"
	"strconv"
)

const createSlot = "create_parking_lot"
const park = "park"
const leaveSlot = "leave"
const status = "status"
const regNumsForColor = "registration_numbers_for_cars_with_colour"
const slotNumsForColor = "slot_numbers_for_cars_with_colour"
const slotNumForRegNum = "slot_number_for_registration_number"

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
