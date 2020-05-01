package main

import (
	"bufio"
	"fmt"
	"os"
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

var slotNum = 0
var slots map[int]carDetails

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		if cmd[0] == "exit" {
			break
		}
		switch cmd[0] {
		case createSlot:
			slotNum, _ = strconv.Atoi(cmd[1])
			if slots == nil {
				slots = make(map[int]carDetails)
			}
		case park:
			regNum := cmd[1]
			color := cmd[2]
			parkCar(regNum, color)
		case leaveSlot:
			slotNum, _ := strconv.Atoi(cmd[1])
			leaveParking(slotNum)
		case status:
			printStatus()
		case regNumsForColor:
			color := cmd[1]
			findRegNumsByColor(color)
		case slotNumsForColor:
			color := cmd[1]
			findSlotNumsByColor(color)
		case slotNumForRegNum:
			regNum := cmd[1]
			findSlotNoByRegNum(regNum)
		}
		fmt.Print(">")
	}
}

func parkCar(regNum, color string) {
	slotsLen := len(slots) + 1
	if slotsLen > slotNum {
		fmt.Println("No slot found!")
		return
	}
	slots[slotsLen] = carDetails{
		regNum: regNum,
		color:  color,
	}
}

func leaveParking(slotNum int) {
	delete(slots, slotNum)
}

func printStatus() {
	fmt.Println("Slot\tRegistration Number\tColor")
	for slot, carDetails := range slots {
		fmt.Printf("%-6v\t%-20s\t%-16s\n", slot, carDetails.regNum, carDetails.color)
	}
}

func findRegNumsByColor(color string) {
	for _, carDetails := range slots {
		if strings.EqualFold(carDetails.color, color) {
			fmt.Print(carDetails.regNum, " ")
		}
	}
	fmt.Println()
}

func findSlotNumsByColor(color string) {
	for slot, carDetails := range slots {
		if strings.EqualFold(carDetails.color, color) {
			fmt.Print(slot, " ")
		}
	}
	fmt.Println()
}

func findSlotNoByRegNum(regNum string) {
	for slot, carDetails := range slots {
		if strings.EqualFold(carDetails.regNum, regNum) {
			fmt.Print(slot, " ")
		}
	}
	fmt.Println()
}
