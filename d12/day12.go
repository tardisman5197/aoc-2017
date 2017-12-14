package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type program struct {
	id    int
	pipes []int
}

func main() {
	data := readFile(os.Args[1])
	fmt.Println(part1(data))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filename string) map[int]program {
	// Open File
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	// Parse file
	data := make(map[int]program, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		row := strings.Fields(s.Text())
		// Create program
		var currentProgram program
		for index, value := range row {
			if index == 0 {
				id, _ := strconv.Atoi(value)
				if _, ok := data[id]; !ok {
					currentProgram.id, _ = strconv.Atoi(value)
				} else {
					currentProgram = data[id]
				}
			} else if index > 1 {
				if value[len(value)-1] == ',' {
					value = value[:len(value)-1]
				}
				currentPipe, _ := strconv.Atoi(value)
				if _, ok := data[currentPipe]; !ok {
					data[currentPipe] = program{id: currentPipe, pipes: []int{}}
				}
				currentProgram.pipes = append(currentProgram.pipes, currentPipe)
			}
		}
		// Add program to data
		data[currentProgram.id] = currentProgram
	}
	return data
}

func inGroup(data *map[int]program, destination, start, parent int) bool {
	if start == destination {
		return true
	}
	for _, value := range (*data)[start].pipes {
		if value != parent {
			if inGroup(data, destination, value, start) {
				return true
			}
		}
	}
	return false
}

func part1(data map[int]program) int {
	var size = 0
	for key := range data {
		if inGroup(&data, 0, key, -1) {
			size++
		}
	}
	return size
}
