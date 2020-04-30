package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const crateSlot = "create_parking_lot"
const parkCar = "park"
const leaveSlot = "leave"
const status = "status"
const regNumsForColor = "registration_numbers_for_cars_with_colour"
const slotNumsForColor = "slot_numbers_for_cars_with_colour"
const slotNumForRegNum = "slot_number_for_registration_number"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		switch cmd[0] {
		case crateSlot:
			slotNums := cmd[1]
			fmt.Println(slotNums)
		case parkCar:
			regNum := cmd[1]
			color := cmd[2]
			fmt.Println(regNum, color)
		case leaveSlot:
			slotNum := cmd[1]
			fmt.Println(slotNum)
		case status:
			fmt.Println(cmd[0])
		case regNumsForColor:
			color := cmd[1]
			fmt.Println(color)
		case slotNumsForColor:
			color := cmd[1]
			fmt.Println(color)
		case slotNumForRegNum:
			regNum := cmd[1]
			fmt.Println(regNum)
		}
	}

}
