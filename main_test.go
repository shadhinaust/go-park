package main

import "testing"

type testSlot struct {
	cmd, res []string
}

var testData = []testSlot{
	{
		cmd: []string{"create_parking_lot 5",
			"park TY90F0 RED",
			"park XY71Q0 BLUE",
			"park WY54P0 GREEN",
			"park UY39F0 RED",
			"leave 2",
			"park TR10F9 WHITE",
			"status",
			"park PY90E5 RED",
			"leave 5",
			"park DT74Y0 YELLOW",
			"registration_numbers_for_cars_with_colour RED",
			"slot_numbers_for_cars_with_colour GREEN",
			"slot_number_for_registration_number TY90F0"},
		res: []string{"", "", "", "", "", "", "", "status", "", "", "", "reg num", "slot nums", "slot num"},
	},
	{
		cmd: []string{"create_parking_lot 5",
			"park TY90F0 RED",
			"park XY71Q0 BLUE",
			"park WY54P0 GREEN",
			"park UY39F0 RED",
			"leave 2",
			"park TR10F9 WHITE",
			"status",
			"park PY90E5 RED",
			"leave 5",
			"park DT74Y0 YELLOW",
			"registration_numbers_for_cars_with_colour RED",
			"slot_numbers_for_cars_with_colour GREEN",
			"slot_number_for_registration_number TY90F0"},
		res: []string{"", "", "", "", "", "", "", "status", "", "", "", "reg num", "slot nums", "slot num"},
	},
}

func TestExecCmd(t *testing.T) {
	for _, test := range testData {
		t.
	}
}
