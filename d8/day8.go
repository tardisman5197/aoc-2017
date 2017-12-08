package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type condition struct {
	register string
	logic    string
	amount   int
}
type instruction struct {
	register  string
	operation string
	amount    int
	condition condition
}

func readFile(filename string) []instruction {
	// Open File
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	// Parse file
	instructions := make([]instruction, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		row := strings.Fields(s.Text())
		var currentInstruction instruction
		for index, value := range row {
			switch index {
			case 0:
				currentInstruction.register = value
			case 1:
				currentInstruction.operation = value
			case 2:
				currentInstruction.amount, _ = strconv.Atoi(value)
			case 4:
				currentInstruction.condition.register = value
			case 5:
				currentInstruction.condition.logic = value
			case 6:
				currentInstruction.condition.amount, _ = strconv.Atoi(value)
			}
		}
		instructions = append(instructions, currentInstruction)
	}
	return instructions
}

func main() {
	instructions := readFile(os.Args[1])
	part1(instructions)

}

func conditionMet(registers *map[string]int, condition condition) bool {
	if _, ok := (*registers)[condition.register]; !ok {
		(*registers)[condition.register] = 0
	}
	switch condition.logic {
	case ">":
		//fmt.Println("one")
	case "<":
		//fmt.Println("two")
	case "<=":
		//fmt.Println("three")
	case ">=":
		//fmt.Println("three")
	case "!=":
		//fmt.Println("three")
	}
	return false
}
func part1(instructions []instruction) int {
	registers := make(map[string]int)
	for _, value := range instructions {
		if conditionMet(&registers, value.condition) {

		}
	}
	return 1
}

func part2() {

}
