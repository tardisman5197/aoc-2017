package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Open File
	f, err := os.Open(os.Args[1])
	check(err)
	defer f.Close()

	// Parse file
	data := make([]int, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		row := strings.Fields(s.Text())
		for _, num := range row {
			val, _ := strconv.Atoi(num)
			data = append(data, val)
		}
	}

	part1(data)
	// 	part2(data)
}

func part1(data []int) {
	past := make([][]int, 0)
	for {
		// Get Lowest
		var highest = 0
		for i, bank := range data {
			if bank > data[highest] {
				highest = i
			}
		}

		// Distribute blocks
		var block = data[highest]
		data[highest] = 0
		for i := 1; i <= block; i++ {
			data[(highest+i)%len(data)]++
		}

		// Check for repeat
		for _, bank := range past {
			if reflect.DeepEqual(data, bank) {
				fmt.Println(len(past) + 1)
				return
			}
		}

		// Add data to past
		dataC := make([]int, len(data))
		copy(dataC, data)
		past = append(past, dataC)
	}
}
