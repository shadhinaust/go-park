package executor

import (
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
		ParkCar(regNum, color)
	case leaveSlot:
		slotNum, _ := strconv.Atoi(cmd[1])
		LeaveParking(slotNum)
	case status:
		PrintStatus()
	case regNumsForColor:
		color := cmd[1]
		FindRegNumsByColor(color)
	case slotNumsForColor:
		color := cmd[1]
		FindSlotNumsByColor(color)
	case slotNumForRegNum:
		regNum := cmd[1]
		FindSlotNoByRegNum(regNum)
	}
}
