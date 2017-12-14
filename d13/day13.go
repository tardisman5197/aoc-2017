package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type scanner struct {
	depth int
	area  int
}

func main() {
	input := readFile(os.Args[1])
	fmt.Println(part1(&input))
	fmt.Println(part2(&input))
}

func readFile(filename string) map[int]scanner {
	// Open File
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Parse file
	data := make(map[int]scanner)
	s := bufio.NewScanner(f)
	for s.Scan() {
		row := strings.Fields(s.Text())
		var currentScanner scanner
		currentScanner.depth, _ = strconv.Atoi(row[0][:len(row[0])-1])
		currentScanner.area, _ = strconv.Atoi(row[1])
		data[currentScanner.depth] = currentScanner
	}
	//fmt.Println(data)
	return data
}

func part1(scanners *map[int]scanner) int {
	var severity = 0
	for _, scanner := range *scanners {
		if scanner.depth%(2*(scanner.area-1)) == 0 {
			severity += scanner.depth * scanner.area
		}
	}
	return severity
}

func isCaught(scanners *map[int]scanner, delay int) bool {
	for _, scanner := range *scanners {
		if (scanner.depth+delay)%(2*(scanner.area-1)) == 0 {
			return true
		}
	}
	return false
}

func part2(scanners *map[int]scanner) int {
	var delay = 0
	for {
		if !isCaught(scanners, delay) {
			break
		}
		delay++
	}
	return delay
}
