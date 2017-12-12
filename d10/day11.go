package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		for _, num := range row {
			val, _ := strconv.Atoi(num)
			data = append(data, val)
		}
	}
	return data
}

func main() {
	input := readFile(os.Args[1])
	size, _ := strconv.Atoi(os.Args[2])
	fmt.Println(part1(size, input))
}

func part1(size int, lengths []int) int {
	// Setup list
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}

	var skip = 0
	var currentIndex = 0
	for _, length := range lengths {
		for i := 0; i < length/2; i++ {
			tmp := list[(currentIndex+i)%len(list)]
			list[(currentIndex+i)%len(list)] = list[((currentIndex + length - 1 - i) % len(list))]
			list[((currentIndex + length - 1 - i) % len(list))] = tmp
		}
		currentIndex = (currentIndex + length + skip) % len(list)
		skip++
	}
	return list[0] * list[1]
}
