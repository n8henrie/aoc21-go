package d01

import (
	"testing"
)

var exampleInput = `199
200
208
210
200
207
240
269
260
263`

func TestParse(t *testing.T) {
	result, err := parseInput(exampleInput)
	if err != nil {
		t.Error(err)
	}
	length := len(result)
	if length != 10 {
		t.Errorf("Wrong length of parsed input, expected %d found %d", 10, length)
	}
}

func TestPackage(t *testing.T) {
	input, _ := parseInput(exampleInput)
	if part1(input) != 7 {
		t.Error("Part 1 failed")
	}
}
