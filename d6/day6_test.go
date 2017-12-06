package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("test.txt")
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
	result, _, _ := part1(data)
	if result != 5 {
		t.Fatalf("Part 1 wrong: %d", result)
	}
}

func TestPart2(t *testing.T) {
	f, err := os.Open("test.txt")
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
	result := part2(data)
	if result != 4 {
		t.Fatalf("Part 1 wrong: %d", result)
	}
}
