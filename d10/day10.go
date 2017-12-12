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

func readFile1(filename string) []int {
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

func readFile2(filename string) []byte {
	// Open File
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	// Parse file
	data := make([]byte, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		row := s.Bytes()
		for _, byteVal := range row {
			data = append(data, byteVal)
		}
	}
	data = append(data, 17, 31, 73, 47, 23)
	return data
}

func main() {
	size, _ := strconv.Atoi(os.Args[2])

	input := readFile1(os.Args[1])
	fmt.Println(part1(size, input))

	inputByte := readFile2(os.Args[1])
	fmt.Println(part2(size, inputByte))
}

func round(list, lengths []int, startIndex, skip int) ([]int, int, int) {
	var currentIndex = startIndex
	for _, length := range lengths {
		for i := 0; i < length/2; i++ {
			tmp := list[(currentIndex+i)%len(list)]
			list[(currentIndex+i)%len(list)] = list[((currentIndex + length - 1 - i) % len(list))]
			list[((currentIndex + length - 1 - i) % len(list))] = tmp
		}
		currentIndex = (currentIndex + length + skip) % len(list)
		skip++
	}
	return list, currentIndex, skip
}

func hash(list []int) []int {
	denseHash := make([]int, 0)
	var currentByte int
	for i := 1; i < len(list)+1; i++ {
		currentByte = currentByte ^ list[i-1]
		if i%16 == 0 {
			denseHash = append(denseHash, currentByte)
			currentByte = 0
		}
	}
	return denseHash
}

func encodeHex(list []int) string {
	var resultstr = ""
	for _, value := range list {
		current := strconv.FormatInt(int64(value), 16)
		if len(current) < 2 {
			resultstr += "0"
		}
		resultstr += current
	}
	return resultstr
}

func part1(size int, lengths []int) int {
	// Setup list
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}
	// Do round
	list, _, _ = round(list, lengths, 0, 0)
	// Return the product of the first 2 elements in the list
	return list[0] * list[1]
}

func part2(size int, lengths []byte) string {
	// Setup list
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}
	// Change lengths to ints
	lengthsint := make([]int, 0)
	for _, value := range lengths {
		lengthsint = append(lengthsint, int(value))
	}
	// Do 64 rounds
	var currentIndex, skip = 0, 0
	for i := 0; i < 64; i++ {
		list, currentIndex, skip = round(list, lengthsint, currentIndex, skip)
	}
	// Hash and convert to hex
	list = hash(list)
	resultStr := encodeHex(list)
	return resultStr
}
