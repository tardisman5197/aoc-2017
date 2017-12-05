package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		row, _ := strconv.Atoi(s.Text())
		data = append(data, row)
	}
	part1(data)
}

func part1(data []int) {
	var jumps, index = 0, 0
	for {
		jumps++
		currentValue := data[index]
		data[index]++
		index += currentValue
		if index >= len(data) || index < 0 {
			break
		}
	}
	fmt.Println(jumps)
}
