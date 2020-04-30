package main

import "fmt"

const crateSlot = "create_parking_lot"
const parkCar = "park"
const leaveSlot = "leave"
const status = "status"
const regNumsForColor = "registration_numbers_for_cars_with_colour"
const slotNumsForColor = "slot_numbers_for_cars_with_colour"
const slotNumForRegNum = "slot_number_for_registration_number"

func main() {
	switch input := "input"; input {
	case crateSlot:
		fmt.Println(input)
	case parkCar:
		fmt.Println(input)
	case leaveSlot:
		fmt.Println(input)
	case status:
		fmt.Println(input)
	case regNumsForColor:
		fmt.Println(input)
	case slotNumsForColor:
		fmt.Println(input)
	case slotNumForRegNum:
		fmt.Println(input)
	default:
		fmt.Println(input)
	}

}
