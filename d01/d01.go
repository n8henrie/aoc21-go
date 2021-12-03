package d01

import (
	"aoc21/utils/files"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int64, error) {
	lines := strings.Split(input, "\n")
	length := len(lines)
	result := make([]int64, length)
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			continue
		}
		asInt, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		result[i] = asInt
	}
	return result, nil
}

func part1(input []int64) int {
	result := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			result++
		}
	}
	return result
}

func part2(input []int64) int {
	sums := make([]int64, len(input)-2)
	for i := 0; i < len(input)-3; i++ {
		sums[i] = input[i] + input[i+1] + input[i+2]
	}
	return part1(sums)
}

func Run() {
	raw, err := files.InputForDay(1)
	if err != nil {
		log.Fatal(err)
	}
	input, err := parseInput(raw)
	if err != nil {
		log.Fatal(err)
	}
	part1Result := part1(input)
	fmt.Printf("Part 1: %d\n", part1Result)
	part2Result := part2(input)
	fmt.Printf("Part 2: %d\n", part2Result)
}
