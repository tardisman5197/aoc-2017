package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filename string) []int {
	// Open File
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	// Parse file
	data := make([]int, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		row := strings.Split(s.Text(), ",")
		for _, value := range row {
			switch value {
			case "n":
				data = append(data, 0)
			case "ne":
				data = append(data, 1)
			case "se":
				data = append(data, 2)
			case "s":
				data = append(data, 3)
			case "sw":
				data = append(data, 4)
			case "nw":
				data = append(data, 5)
			}
		}
	}
	return data
}

func main() {
	data := readFile(os.Args[1])
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func part1(data []int) int {
	// Init directions map
	directions := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
	// Add data to directions
	for _, direction := range data {
		directions[direction]++
	}
	var distance = 0
	// Remove opposits
	for i := 0; i < len(directions)/2; i++ {
		if directions[i] > directions[i+3] {
			directions[i] -= directions[i+3]
			directions[i+3] = 0
		} else {
			directions[i+3] -= directions[i]
			directions[i] = 0
		}
	}
	// Condense directions
	for i := 0; i < len(directions)/2; i++ {
		if directions[i] > directions[i+2] {
			directions[i+1] += directions[i+2]
			directions[i] -= directions[i+2]
			directions[i+2] = 0
		} else {
			directions[i+1] += directions[i]
			directions[i+2] -= directions[i]
			directions[i] = 0
		}
	}
	// Get distance from start
	for _, value := range directions {
		distance += value
	}
	return distance
}

func part2(data []int) int {
	// Init coords
	var x, y, z, highestDistance float64 = 0, 0, 0, 0
	// Loop though each direction and update coords
	for _, value := range data {
		switch value {
		case 0:
			z++
			x++
		case 1:
			x++
			y--
		case 2:
			y--
			z++
		case 3:
			z--
			x--
		case 4:
			y++
			x--
		case 5:
			y++
			z--
		}
		// Check if further away
		if highestDistance < math.Abs(x) {
			highestDistance = math.Abs(x)
		}
		if highestDistance < math.Abs(y) {
			highestDistance = math.Abs(y)
		}
		if highestDistance < math.Abs(z) {
			highestDistance = math.Abs(z)
		}
	}
	return int(highestDistance)
}
