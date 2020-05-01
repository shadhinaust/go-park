package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		if cmd[0] == "exit" {
			return
		}
		switch cmd[0] {
		case createSlot:
			slotNum, _ := strconv.Atoi(cmd[1])
			createParkingSlots(slotNum)
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

func createParkingSlots(slotNum int) {
	slots = make([]carDetails, slotNum)
}

func parkCar(regNum, color string) {
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

func leaveParking(slotNum int) {
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

func printStatus() {
	fmt.Println("Slot\tRegistration Number\tColor")
	for slot, details := range slots {
		if details != (carDetails{}) {
			fmt.Printf("%-6v\t%-20s\t%-16s\n", slot+1, details.regNum, details.color)
		}
	}
}

func findRegNumsByColor(color string) {
	for _, details := range slots {
		if strings.EqualFold(details.color, color) {
			fmt.Printf("%s\n", details.regNum)
		}
	}
}

func findSlotNumsByColor(color string) {
	for slot, details := range slots {
		if strings.EqualFold(details.color, color) {
			fmt.Printf("%d\n", slot+1)
		}
	}
}

func findSlotNoByRegNum(regNum string) {
	for slot, details := range slots {
		if strings.EqualFold(details.regNum, regNum) {
			fmt.Printf("%d\n", slot+1)
		}
	}
}
